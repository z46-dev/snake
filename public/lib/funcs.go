package lib

import "syscall/js"

func Now() float64 {
	return js.Global().Get("performance").Call("now").Float()
}

func UiScale(c *CanvasStruct) float64 {
	width, height := c.Size()

	if height > width {
		return height / 1080.0
	}

	return width / 1920.0
}

func SetupRenderingFrame(c *CanvasStruct, function func(float64, float64)) {
	c.Resize()
	var Width, Height = c.Size()

	var frame js.Func
	frame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		function(Width, Height)
		js.Global().Call("requestAnimationFrame", frame)
		return nil
	})

	js.Global().Call("requestAnimationFrame", frame)

	// Resize on window resize
	js.Global().Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		c.Resize()
		Width, Height = c.Size()
		return nil
	}))
}
