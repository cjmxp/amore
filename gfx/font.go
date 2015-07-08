package gfx

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/go-gl/gl/v2.1/gl"

	"github.com/tanema/freetype-go/freetype"
)

var current_font *Font

type Font struct {
	Texture *Texture
	Offset  int
	Glyphs  map[rune]Glyph
}

type textRec struct {
	X1, Y1, X2, Y2 float32
}

type Glyph struct {
	TextureRec    textRec
	Width, Height int
	Advance       int
}

func pow2(x uint32) uint32 {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	return x + 1
}

func NewFont(filename string, font_size float32) (*Font, error) {
	fontBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	ttf, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	glyphs := ttf.ListRunes()
	glyph_dict := make(map[rune]Glyph)
	glyphsPerRow := int32(16)
	glyphsPerCol := (int32(len(glyphs)) / glyphsPerRow) + 1

	context := freetype.NewContext()
	context.SetDPI(72)
	context.SetFont(ttf)
	context.SetFontSize(float64(font_size))

	font_bounds := ttf.Bounds()
	glyph_width := context.FUnitToPixelRU(int(font_bounds.XMax - font_bounds.XMin))
	glyph_height := context.FUnitToPixelRU(int(font_bounds.YMax-font_bounds.YMin) + 5)
	image_width := pow2(uint32(int32(glyph_width) * glyphsPerRow))
	image_height := pow2(uint32(int32(glyph_height) * glyphsPerCol))

	rgba := image.NewRGBA(image.Rect(0, 0, int(image_width), int(image_height)))
	context.SetClip(rgba.Bounds())
	context.SetDst(rgba)
	context.SetSrc(image.White)

	var gx, gy int
	offset := context.FUnitToPixelRU(ttf.UnitsPerEm())
	for i, ch := range glyphs {
		pt := freetype.Pt(gx+offset, gy+offset)
		context.DrawString(string(ch), pt)

		tx1 := float32(gx) / float32(image_width)
		ty1 := float32(gy) / float32(image_height)
		tx2 := (float32(gx) + float32(glyph_width)) / float32(image_width)
		ty2 := (float32(gy) + float32(glyph_height)) / float32(image_height)

		index := ttf.Index(ch)
		metric := ttf.HMetric(index)
		glyph_dict[ch] = Glyph{
			TextureRec: textRec{tx1, ty1, tx2, ty2},
			Width:      glyph_width,
			Height:     glyph_height,
			Advance:    int(context.FUnitToPixelRU(int(metric.AdvanceWidth))),
		}

		if i%16 == 0 {
			gx = 0
			gy += glyph_height
		} else {
			gx += glyph_width
		}
	}

	texture, err := NewImageTexture(rgba)
	if err != nil {
		return nil, err
	}

	return &Font{
		Texture: texture,
		Glyphs:  glyph_dict,
		Offset:  offset,
	}, nil
}

func NewImageFont(filename, glyph_hints string) (*Font, error) {
	glyph_rune_hints := []rune(glyph_hints)

	imgFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	glyph_dict := make(map[rune]Glyph)
	glyphsPerRow := int32(16)
	glyphsPerCol := (int32(len(glyph_rune_hints)) / glyphsPerRow) + 1
	glyph_width := img.Bounds().Dx() / len(glyph_rune_hints)
	glyph_height := img.Bounds().Dy()
	image_width := pow2(uint32(int32(glyph_width) * glyphsPerRow))
	image_height := pow2(uint32(int32(glyph_height) * glyphsPerCol))

	rgba := image.NewRGBA(image.Rect(0, 0, int(image_width), int(image_height)))

	var gx, gy int
	for i := 0; i < len(glyph_rune_hints); i++ {
		dst_rec := image.Rect(gx, gy, gx+glyph_width, gy+glyph_height)
		draw.Draw(rgba, dst_rec, img, image.Pt(i*glyph_width, 0), draw.Src)

		tx1 := float32(gx) / float32(image_width)
		ty1 := float32(gy) / float32(image_height)
		tx2 := (float32(gx) + float32(glyph_width)) / float32(image_width)
		ty2 := (float32(gy) + float32(glyph_height)) / float32(image_height)

		glyph_dict[glyph_rune_hints[i]] = Glyph{
			TextureRec: textRec{tx1, ty1, tx2, ty2},
			Width:      glyph_width,
			Height:     glyph_height,
			Advance:    glyph_width,
		}

		if i%16 == 0 {
			gx = 0
			gy += glyph_height
		} else {
			gx += glyph_width
		}
	}

	texture, err := NewImageTexture(rgba)
	if err != nil {
		return nil, err
	}

	return &Font{
		Texture: texture,
		Glyphs:  glyph_dict,
	}, nil
}

func (f *Font) Index(ch rune) {
}

func (f *Font) Release() {
	f.Texture.Release()
}

func SetFont(font *Font) {
	current_font = font
}

func Printf(x, y float32, fs string, argv ...interface{}) {
	if current_font == nil {
		return
	}

	formatted_string := fmt.Sprintf(fs, argv...)
	if len(formatted_string) == 0 {
		return
	}

	x = x - float32(current_font.Offset)
	y = y - (float32(current_font.Offset) / 4.0)

	current_font.Texture.Bind(func() {
		for _, ch := range formatted_string {
			if glyph, ok := current_font.Glyphs[ch]; ok {
				gl.Begin(gl.QUADS)
				{
					gl.TexCoord2f(glyph.TextureRec.X1, glyph.TextureRec.Y1) // top-left
					gl.Vertex2f(x, y)
					gl.TexCoord2f(glyph.TextureRec.X1, glyph.TextureRec.Y2) // bottom-left
					gl.Vertex2f(x, y+float32(glyph.Height))
					gl.TexCoord2f(glyph.TextureRec.X2, glyph.TextureRec.Y2) // bottom-right
					gl.Vertex2f(x+float32(glyph.Width), y+float32(glyph.Height))
					gl.TexCoord2f(glyph.TextureRec.X2, glyph.TextureRec.Y1) // top-right
					gl.Vertex2f(x+float32(glyph.Width), y)
				}
				gl.End()
				x = x + float32(glyph.Advance)
			}
		}
	})

}