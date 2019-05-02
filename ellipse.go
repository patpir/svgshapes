package svgshapes

import (
	"encoding/xml"
)

type Ellipse struct {
	XMLName xml.Name `xml:"ellipse"`
	Shapebase
	CenterX float64 `xml:"cx,attr"`
	CenterY float64 `xml:"cy,attr"`
	RadiusX float64 `xml:"rx,attr"`
	RadiusY float64 `xml:"ry,attr"`
}

