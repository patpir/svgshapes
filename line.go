package svgshapes

import (
	"encoding/xml"
)

type Line struct {
	XMLName xml.Name `xml:"line"`
	Shapebase
	StartX float64 `xml:"x1,attr"`
	StartY float64 `xml:"y1,attr"`
	EndX   float64 `xml:"x2,attr"`
	EndY   float64 `xml:"y2,attr"`
}

