package lib

import "syscall/js"

type Path2D struct {
	Path2DValue js.Value
}

func NewPath2D() *Path2D {
	return &Path2D{js.Global().Get("Path2D").New()}
}

func NewPath2DFromString(str string) *Path2D {
	return &Path2D{js.Global().Get("Path2D").New(str)}
}

func (p *Path2D) MoveTo(x, y float64) {
	p.Path2DValue.Call("moveTo", x, y)
}

func (p *Path2D) LineTo(x, y float64) {
	p.Path2DValue.Call("lineTo", x, y)
}

func (p *Path2D) Arc(x, y, radius, startAngle, endAngle float64) {
	p.Path2DValue.Call("arc", x, y, radius, startAngle, endAngle)
}

func (p *Path2D) ArcTo(x1, y1, x2, y2, radius float64) {
	p.Path2DValue.Call("arcTo", x1, y1, x2, y2, radius)
}

func (p *Path2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	p.Path2DValue.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

func (p *Path2D) QuadraticCurveTo(cpx, cpy, x, y float64) {
	p.Path2DValue.Call("quadraticCurveTo", cpx, cpy, x, y)
}

func (p *Path2D) Rect(x, y, width, height float64) {
	p.Path2DValue.Call("rect", x, y, width, height)
}

func (p *Path2D) ClosePath() {
	p.Path2DValue.Call("closePath")
}
