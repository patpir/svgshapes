package svgshapes

import (
	"encoding/xml"
	"strconv"
)

type Path struct {
	XMLName xml.Name `xml:"path"`
	Shapebase
	PathCommands string `xml:"d,attr"`
}

func (p *Path) addCommand(command string, parameters ...float64) {
	cmdString := command
	for _, param := range parameters {
		cmdString += " " + strconv.FormatFloat(param, 'f', -1, 64)
	}

	if len(p.PathCommands) == 0 {
		p.PathCommands = cmdString
	} else {
		p.PathCommands = p.PathCommands + " " + cmdString
	}
}

func (p *Path) MoveTo(x, y float64) *Path {
	p.addCommand("M", x, y)
	return p
}

func (p *Path) MoveRelative(dx, dy float64) *Path {
	p.addCommand("m", dx, dy)
	return p
}

func (p *Path) ClosePath() *Path {
	p.addCommand("Z")
	return p
}

func (p *Path) LineTo(x, y float64) *Path {
	p.addCommand("L", x, y)
	return p
}

func (p *Path) LineRelative(dx, dy float64) *Path {
	p.addCommand("l", dx, dy)
	return p
}

func (p *Path) HLineTo(x float64) *Path {
	p.addCommand("H", x)
	return p
}

func (p *Path) HLineRelative(dx float64) *Path {
	p.addCommand("h", dx)
	return p
}

func (p *Path) VLineTo(y float64) *Path {
	p.addCommand("V", y)
	return p
}

func (p *Path) VLineRelative(dy float64) *Path {
	p.addCommand("v", dy)
	return p
}

func (p *Path) CubicBezierTo(controlX1, controlY1, controlX2, controlY2, x, y float64) *Path {
	p.addCommand("C", controlX1, controlY1, controlX2, controlY2, x, y)
	return p
}

func (p *Path) CubicBezierRelative(controlDX1, controlDY1, controlDX2, controlDY2, dx, dy float64) *Path {
	p.addCommand("c", controlDX1, controlDY1, controlDX2, controlDY2, dx, dy)
	return p
}

func (p *Path) SmoothCubicBezierTo(controlX2, controlY2, x, y float64) *Path {
	p.addCommand("S", controlX2, controlY2, x, y)
	return p
}

func (p *Path) SmoothCubicBezierRelative(controlDX2, controlDY2, dx, dy float64) *Path {
	p.addCommand("s", controlDX2, controlDY2, dx, dy)
	return p
}

func (p *Path) QuadraticBezierTo(controlX, controlY, x, y float64) *Path {
	p.addCommand("Q", controlX, controlY, x, y)
	return p
}

func (p *Path) QuadraticBezierRelative(controlDX, controlDY, dx, dy float64) *Path {
	p.addCommand("q", controlDX, controlDY, dx, dy)
	return p
}

func (p *Path) SmoothQuadraticBezierTo(x, y float64) *Path {
	p.addCommand("T", x, y)
	return p
}

func (p *Path) SmoothQuadraticBezierRelative(dx, dy float64) *Path {
	p.addCommand("t", dx, dy)
	return p
}

func (p *Path) addEllipticalArcCommand(command string, radiusX, radiusY, xAxisRotation float64, largeArc, sweep bool, x, y float64) {
	largeArcNum := 0.0
	if largeArc {
		largeArcNum = 1.0
	}
	sweepNum := 0.0
	if sweep {
		sweepNum = 1.0
	}
	p.addCommand(command, radiusX, radiusY, xAxisRotation, largeArcNum, sweepNum, x, y)
}

func (p *Path) EllipticalArcTo(radiusX, radiusY, xAxisRotation float64, largeArc, sweep bool, x, y float64) *Path {
	p.addEllipticalArcCommand("A", radiusX, radiusY, xAxisRotation, largeArc, sweep, x, y)
	return p
}

func (p *Path) EllipticalArcRelative(radiusX, radiusY, xAxisRotation float64, largeArc, sweep bool, dx, dy float64) *Path {
	p.addEllipticalArcCommand("a", radiusX, radiusY, xAxisRotation, largeArc, sweep, dx, dy)
	return p
}

