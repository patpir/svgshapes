package svgshapes

import (
	"encoding/xml"
)

type Circle struct {
	XMLName xml.Name `xml:"circle"`
	Shapebase
	CenterX float64 `xml:"cx,attr"`
	CenterY float64 `xml:"cy,attr"`
	Radius  float64 `xml:"r,attr"`
}

