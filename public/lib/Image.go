package lib

import "syscall/js"

type Image struct {
	ImageValue    js.Value
	Width, Height float64
	Ready         bool
}

func NewImage(src string) *Image {
	var image *Image = new(Image)
	image.ImageValue = js.Global().Get("Image").New()
	image.ImageValue.Set("src", src)
	image.ImageValue.Set("onload", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		image.Width = image.ImageValue.Get("width").Float()
		image.Height = image.ImageValue.Get("height").Float()
		image.Ready = true
		return nil
	}))

	return image
}
