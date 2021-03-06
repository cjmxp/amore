package main

import (
	"os"
	"text/template"
)

var configTemplate, _ = template.New("config").Parse(`title = "{{.Name}}"				# The window title (string)
identity = ""             # The name of the save directory (string)
icon = ""                 # Filepath to an image to use as the window's icon (string)
width = 800               # The window width (number)
height = 600              # The window height (number)
borderless = false        # Remove all border visuals from the window (boolean)
resizable = true          # Let the window be user-resizable (boolean)
minwidth = 1              # Minimum window width if the window is resizable (number)
minheight = 1             # Minimum window height if the window is resizable (number)
fullscreen = false        # Enable fullscreen (boolean)
fullscreentype = "normal" # Standard fullscreen or desktop fullscreen mode (string)
vsync = true              # Enable vertical sync (boolean)
msaa = 16                 # The number of samples to use with multi-sampled antialiasing (number)
display = 1               # Index of the monitor to show the window in (number)
highdpi = false           # Enable high-dpi mode for the window on a Retina display (boolean)
srgb = false              # Enable sRGB gamma correction when drawing to the screen (boolean)
centered = true           # Center the window in the display
x = -1                    # The x-coordinate of the window's position in the specified display (number)
y = -1                    # The y-coordinate of the window's position in the specified display (number)
`)

var mainTemplate, _ = template.New("main").Parse(`package {{.PackageName}}

import (
	"github.com/tanema/amore"
	"github.com/tanema/amore/gfx"
)

func main() {
	amore.Start(update, draw)
}

func update(deltaTime float32) {
}

func draw() {
	gfx.Print("Hello World",50, 50)
}`)

var bundleTemplate, _ = template.New("bundle").Parse(`package {{.PackageName}}
import (
		"github.com/tanema/amore/file"
)

func init() {
	data := "{{.Data}}"
	file.Register(data)
}
`)

func writeOutTemplate(path string, tmpl *template.Template, data interface{}) error {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}
	return tmpl.Execute(file, data)
}
