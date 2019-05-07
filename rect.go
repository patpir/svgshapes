package svgshapes

import (
	"encoding/xml"
)

type Rect struct {
	XMLName xml.Name `xml:"rect"`
	Shapebase
	StartX  float64 `xml:"x,attr"`
	StartY  float64 `xml:"y,attr"`
	Width   float64 `xml:"width,attr"`
	Height  float64 `xml:"height,attr"`
	RadiusX float64 `xml:"rx,attr,omitempty"`
	RadiusY float64 `xml:"ry,attr,omitempty"`
}

