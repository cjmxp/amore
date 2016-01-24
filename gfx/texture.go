package gfx

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var filters = map[string]int32{"linear": gl.LINEAR, "nearest": gl.NEAREST}
var wraps = map[string]int32{"clamp": gl.CLAMP_TO_EDGE, "repeat": gl.REPEAT}

type (
	WrapMode   int
	FilterMode int
	Filter     struct {
		min, mag, mipmap FilterMode
		anisotropy       float32
	}
	Wrap struct {
		s, t WrapMode
	}
	Texture struct {
		textureId     uint32
		Width, Height float32
		coords        []float32
		texcoords     []float32
		filter        Filter
		wrap          Wrap
		mipmaps       bool
		sRGB          bool
	}
)

func newFilter() Filter {
	return Filter{
		min:        FILTER_LINEAR,
		mag:        FILTER_LINEAR,
		mipmap:     FILTER_NONE,
		anisotropy: 1.0,
	}
}

func newWrap() Wrap {
	return Wrap{s: WRAP_CLAMP, t: WRAP_CLAMP}
}

const (
	WRAP_CLAMP           WrapMode   = WrapMode(gl.CLAMP)
	WRAP_REPEAT          WrapMode   = WrapMode(gl.REPEAT)
	WRAP_MIRRORED_REPEAT WrapMode   = WrapMode(gl.MIRRORED_REPEAT)
	FILTER_NONE          FilterMode = FilterMode(gl.NONE)
	FILTER_LINEAR        FilterMode = FilterMode(gl.LINEAR)
	FILTER_NEAREST       FilterMode = FilterMode(gl.NEAREST)
)

func LoadImageTexture(img image.Image) (*Texture, error) {
	var texture_id uint32
	gl.GenTextures(1, &texture_id)

	bounds := img.Bounds()
	new_texture := &Texture{
		textureId: texture_id,
		Width:     float32(bounds.Dx()),
		Height:    float32(bounds.Dy()),
		wrap:      newWrap(),
		filter:    newFilter(),
	}

	//generate a uniform image and upload to vram
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, bounds, img, image.Point{0, 0}, draw.Src)

	bindTexture(new_texture.GetHandle())
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(bounds.Dx()), int32(bounds.Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	new_texture.SetFilter(FILTER_NEAREST, FILTER_NEAREST)
	new_texture.SetWrap(WRAP_CLAMP, WRAP_CLAMP)

	return new_texture, nil
}

func (texture *Texture) GetHandle() uint32 {
	return texture.textureId
}

func (texture *Texture) generateVerticies() {
	w := texture.GetWidth()
	h := texture.GetHeight()
	texture.coords = []float32{0, 0, 0, h, w, h, w, 0}
	texture.texcoords = []float32{0, 0, 0, 1, 1, 1, 1, 0}
}

func (texture *Texture) SetWrap(wrap_s, wrap_t WrapMode) {
	texture.wrap.s = wrap_s
	texture.wrap.t = wrap_t
	bindTexture(texture.GetHandle())
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, int32(wrap_s))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, int32(wrap_t))
}

func (texture *Texture) SetFilter(min, mag FilterMode) error {
	if !texture.validateFilter() {
		if texture.filter.mipmap != FILTER_NONE && !texture.mipmaps {
			return fmt.Errorf("Non-mipmapped image cannot have mipmap filtering.")
		} else {
			return fmt.Errorf("Invalid texture filter.")
		}
	}

	texture.filter.mag = mag
	texture.filter.min = min
	bindTexture(texture.GetHandle())
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int32(min))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int32(mag))
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAX_ANISOTROPY_EXT, texture.filter.anisotropy)
	return nil
}

func (texture *Texture) validateFilter() bool {
	if !texture.mipmaps && texture.filter.mipmap != FILTER_NONE {
		return false
	}

	if texture.filter.mag != FILTER_LINEAR && texture.filter.mag != FILTER_NEAREST {
		return false
	}

	if texture.filter.min != FILTER_LINEAR && texture.filter.min != FILTER_NEAREST {
		return false
	}

	if texture.filter.mipmap != FILTER_LINEAR && texture.filter.mipmap != FILTER_NEAREST && texture.filter.mipmap != FILTER_NONE {
		return false
	}

	return true
}

func (texture *Texture) GetWidth() float32 {
	return texture.Width
}

func (texture *Texture) GetHeight() float32 {
	return texture.Height
}

func (texture *Texture) Release() {
	deleteTexture(texture.textureId)
}

func (texture *Texture) UnloadVolatile() {
	if texture != nil {
		return
	}
	texture.Release()
	texture = nil
}

func (texture *Texture) drawv(model *mgl32.Mat4, coords, texcoords []float32) {
	prepareDraw(model)
	bindTexture(texture.GetHandle())

	gl.EnableVertexAttribArray(ATTRIB_POS)
	gl.EnableVertexAttribArray(ATTRIB_TEXCOORD)

	gl.VertexAttribPointer(ATTRIB_POS, 2, gl.FLOAT, false, 0, gl.Ptr(coords))
	gl.VertexAttribPointer(ATTRIB_TEXCOORD, 2, gl.FLOAT, false, 0, gl.Ptr(texcoords))

	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)

	gl.DisableVertexAttribArray(ATTRIB_TEXCOORD)
	gl.DisableVertexAttribArray(ATTRIB_POS)
}

func (texture *Texture) Draw(args ...float32) {
	texture.drawv(generateModelMatFromArgs(args), texture.coords, texture.texcoords)
}

func (texture *Texture) Drawq(quad Quad, args ...float32) {
	coords, texcoords := quad.getVertices()
	texture.drawv(generateModelMatFromArgs(args), coords, texcoords)
}
