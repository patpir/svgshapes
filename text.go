package svgshapes

import (
	"encoding/xml"
)

type Text struct {
	XMLName xml.Name `xml:"text"`
	Shapebase
	StartX float64 `xml:"x,attr"`
	StartY float64 `xml:"y,attr"`
	Text   string  `xml:",innerxml"`
}
