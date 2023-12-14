package lib

import "syscall/js"

type CanvasStruct struct {
	CanvasValue   js.Value
	Ctx           js.Value
	Width, Height float64
}

func NewCanvas(inputCanvas js.Value) *CanvasStruct {
	if inputCanvas.IsNull() {
		inputCanvas = js.Global().Get("document").Call("createElement", "canvas")
	}

	var canvas *CanvasStruct = new(CanvasStruct)
	canvas.CanvasValue = inputCanvas
	canvas.Ctx = canvas.CanvasValue.Call("getContext", "2d")
	var devicePixelRatio float64 = js.Global().Get("window").Get("devicePixelRatio").Float()
	canvas.CanvasValue.Set("width", js.Global().Get("innerWidth").Float()*devicePixelRatio)
	canvas.CanvasValue.Set("height", js.Global().Get("innerHeight").Float()*devicePixelRatio)
	canvas.Width = canvas.CanvasValue.Get("width").Float()
	canvas.Height = canvas.CanvasValue.Get("height").Float()

	return canvas
}

func (canvas *CanvasStruct) ResizeValues() {
	canvas.Width = canvas.CanvasValue.Get("width").Float()
	canvas.Height = canvas.CanvasValue.Get("height").Float()
}

func (canvas *CanvasStruct) Resize() {
	var devicePixelRatio float64 = js.Global().Get("window").Get("devicePixelRatio").Float()
	canvas.CanvasValue.Set("width", js.Global().Get("innerWidth").Float()*devicePixelRatio)
	canvas.CanvasValue.Set("height", js.Global().Get("innerHeight").Float()*devicePixelRatio)

	canvas.ResizeValues()
	canvas.SetTextBaseline("middle")
	canvas.SetLineJoin("round")
	canvas.SetLineCap("round")
	canvas.CanvasValue.Call("focus")
}

func (canvas *CanvasStruct) Size() (float64, float64) {
	return canvas.Width, canvas.Height
}

func (canvas *CanvasStruct) SetFill(color string) {
	canvas.Ctx.Set("fillStyle", color)
}

func (canvas *CanvasStruct) SetStroke(color string) {
	canvas.Ctx.Set("strokeStyle", color)
}

func (canvas *CanvasStruct) SetLineWidth(width float64) {
	canvas.Ctx.Set("lineWidth", width)
}

func (canvas *CanvasStruct) SetAlpha(alpha float64) {
	canvas.Ctx.Set("globalAlpha", alpha)
}

func (canvas *CanvasStruct) SetFont(font string) {
	canvas.Ctx.Set("font", font)
}

func (canvas *CanvasStruct) SetTextAlign(align string) {
	canvas.Ctx.Set("textAlign", align)
}

func (canvas *CanvasStruct) SetTextBaseline(baseline string) {
	canvas.Ctx.Set("textBaseline", baseline)
}

func (canvas *CanvasStruct) SetLineCap(cap string) {
	canvas.Ctx.Set("lineCap", cap)
}

func (canvas *CanvasStruct) SetLineJoin(join string) {
	canvas.Ctx.Set("lineJoin", join)
}

func (canvas *CanvasStruct) FillRect(x, y, width, height float64) {
	canvas.Ctx.Call("fillRect", x, y, width, height)
}

func (canvas *CanvasStruct) StrokeRect(x, y, width, height float64) {
	canvas.Ctx.Call("strokeRect", x, y, width, height)
}

func (canvas *CanvasStruct) ClearRect(x, y, width, height float64) {
	canvas.Ctx.Call("clearRect", x, y, width, height)
}

func (canvas *CanvasStruct) FillText(text string, x, y float64) {
	canvas.Ctx.Call("fillText", text, x, y)
}

func (canvas *CanvasStruct) StrokeText(text string, x, y float64) {
	canvas.Ctx.Call("strokeText", text, x, y)
}

func (canvas *CanvasStruct) MeasureText(text string) float64 {
	return canvas.Ctx.Call("measureText", text).Get("width").Float()
}

func (canvas *CanvasStruct) BeginPath() {
	canvas.Ctx.Call("beginPath")
}

func (canvas *CanvasStruct) ClosePath() {
	canvas.Ctx.Call("closePath")
}

func (canvas *CanvasStruct) MoveTo(x, y float64) {
	canvas.Ctx.Call("moveTo", x, y)
}

func (canvas *CanvasStruct) LineTo(x, y float64) {
	canvas.Ctx.Call("lineTo", x, y)
}

func (canvas *CanvasStruct) Arc(x, y, radius, startAngle, endAngle float64) {
	canvas.Ctx.Call("arc", x, y, radius, startAngle, endAngle)
}

func (canvas *CanvasStruct) ArcTo(x1, y1, x2, y2, radius float64) {
	canvas.Ctx.Call("arcTo", x1, y1, x2, y2, radius)
}

func (canvas *CanvasStruct) Ellipse(x, y, radiusX, radiusY, rotation, startAngle, endAngle float64) {
	canvas.Ctx.Call("ellipse", x, y, radiusX, radiusY, rotation, startAngle, endAngle)
}

func (canvas *CanvasStruct) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	canvas.Ctx.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

func (canvas *CanvasStruct) QuadraticCurveTo(cpx, cpy, x, y float64) {
	canvas.Ctx.Call("quadraticCurveTo", cpx, cpy, x, y)
}

func (canvas *CanvasStruct) Rect(x, y, width, height float64) {
	canvas.Ctx.Call("rect", x, y, width, height)
}

func (canvas *CanvasStruct) Fill() {
	canvas.Ctx.Call("fill")
}

func (canvas *CanvasStruct) Stroke() {
	canvas.Ctx.Call("stroke")
}

func (canvas *CanvasStruct) Clip() {
	canvas.Ctx.Call("clip")
}

func (canvas *CanvasStruct) Save() {
	canvas.Ctx.Call("save")
}

func (canvas *CanvasStruct) Restore() {
	canvas.Ctx.Call("restore")
}

func (canvas *CanvasStruct) Translate(x, y float64) {
	canvas.Ctx.Call("translate", x, y)
}

func (canvas *CanvasStruct) Scale(x, y float64) {
	canvas.Ctx.Call("scale", x, y)
}

func (canvas *CanvasStruct) Rotate(angle float64) {
	canvas.Ctx.Call("rotate", angle)
}

func (canvas *CanvasStruct) DrawCanvas(canvas2 *CanvasStruct, x, y float64) {
	canvas.Ctx.Call("drawImage", canvas2.CanvasValue, x, y)
}

func (canvas *CanvasStruct) DrawCanvas2(canvas2 *CanvasStruct, x, y, width, height float64) {
	canvas.Ctx.Call("drawImage", canvas2.CanvasValue, x, y, width, height)
}

func (canvas *CanvasStruct) DrawImage(image *Image, x, y float64) {
	canvas.Ctx.Call("drawImage", image.ImageValue, x, y)
}

func (canvas *CanvasStruct) DrawImage2(image *Image, x, y, width, height float64) {
	canvas.Ctx.Call("drawImage", image.ImageValue, x, y, width, height)
}

func (canvas *CanvasStruct) DrawImage3(image *Image, sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight float64) {
	canvas.Ctx.Call("drawImage", image.ImageValue, sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight)
}

func (canvas *CanvasStruct) FillPath2D(path *Path2D, fill, rule string) {
	canvas.SetFill(fill)
	canvas.Ctx.Call("fill", path.Path2DValue, rule)
}

func (canvas *CanvasStruct) FillPath2DNoRule(path *Path2D, fill string) {
	canvas.SetFill(fill)
	canvas.Ctx.Call("fill", path.Path2DValue)
}

func (canvas *CanvasStruct) StrokePath2D(path *Path2D, stroke string, lineWidth float64) {
	canvas.SetStroke(stroke)
	canvas.SetLineWidth(lineWidth)
	canvas.Ctx.Call("stroke", path.Path2DValue)
}

func (canvas *CanvasStruct) ClipPath2D(path *Path2D, rule string) {
	canvas.Ctx.Call("clip", path.Path2DValue, rule)
}
