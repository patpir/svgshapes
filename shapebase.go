package svgshapes

import ()

type Shapebase struct {
	Id     string `xml:"id,attr,omitempty"`
	Class  string `xml:"class,attr,omitempty"`
	Style  string `xml:"style,attr,omitempty"`
	Fill   string `xml:"fill,attr,omitempty"`
	Stroke string `xml:"stroke,attr,omitempty"`
}

