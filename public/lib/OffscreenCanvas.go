package lib

import "syscall/js"

type OffscreenCanvas struct {
	CanvasStruct
}

func NewOffscreenCanvas(width, height float64) *OffscreenCanvas {
	var canvas *OffscreenCanvas = new(OffscreenCanvas)
	canvas.CanvasValue = js.Global().Get("document").Call("createElement", "canvas")
	canvas.CanvasValue.Set("width", width)
	canvas.CanvasValue.Set("height", height)
	canvas.Ctx = canvas.CanvasValue.Call("getContext", "2d")
	canvas.Width = width
	canvas.Height = height

	return canvas
}
