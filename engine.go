// Copywrite 2016 Tim Anema. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be foind in the LICENSE file.

/*
The base amore package is simply for stopping and starting your game. It will
also automatically lock the os thread and set the application to use all available
cpus.
*/
package amore

import (
	"runtime"

	"github.com/tanema/amore/event"
	"github.com/tanema/amore/gfx"
	"github.com/tanema/amore/timer"
	"github.com/tanema/amore/window"
)

var (
	current_window *window.Window
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.LockOSThread() //important SDL and OpenGl Demand it and stamp thier feet if you dont
}

// Start creates a window and context for the game to run on and runs the game
// loop. As such this function should be put as the last call in your main function.
// update and draw will be called synchronously because calls to OpenGL that are
// not on the main thread will crash your program.
func Start(update func(float32), draw func()) {
	current_window = window.GetCurrent()
	defer current_window.Destroy()
	gfx.InitContext(current_window.GetWidth(), current_window.GetHeight())
	defer gfx.DeInit()
	for !current_window.ShouldClose() {
		timer.Step()
		update(timer.GetDelta())
		gfx.ClearC(gfx.GetBackgroundColorC())
		gfx.Origin()
		draw()
		gfx.Present()
		event.Poll()
	}
}

// Quit will prepare the window to close at the end of the next game loop. This
// will allow a nice clean destruction of all object that are allocated in OpenGL,
// SDL, and OpenAl
func Quit() {
	current_window.SetShouldClose(true)
}
