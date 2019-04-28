package svgshapes

import (
	"encoding/xml"
)

type Polygon struct {
	XMLName xml.Name `xml:"polygon"`
	Shapebase
	Points PointList `xml:"points,attr"`
}

