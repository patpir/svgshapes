package svgshapes

import (
	"encoding/xml"
)

type Polyline struct {
	XMLName xml.Name `xml:"polyline"`
	Shapebase
	Points PointList `xml:"points,attr"`
}

