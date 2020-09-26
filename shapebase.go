package svgshapes

import (
	"encoding/xml"
)

type Shapebase struct {
	Id         string `xml:"id,attr,omitempty"`
	Class      string `xml:"class,attr,omitempty"`
	Style      string `xml:"style,attr,omitempty"`
	Fill       string `xml:"fill,attr,omitempty"`
	Stroke     string `xml:"stroke,attr,omitempty"`
	Attributes []Attribute `xml:",attr"`
}

func (s *Shapebase) SetAttribute(name, value string) {
	attr := rawAttribute{
		Name: xml.Name{
			Space: "",
			Local: name,
		},
		Value: value,
	}
	replacedExistingAttr := false
	for i, existingAttr := range s.Attributes {
		if existingAttr.XMLName() == attr.Name {
			s.Attributes[i] = attr
			replacedExistingAttr = true
		}
	}
	if !replacedExistingAttr {
		s.Attributes = append(s.Attributes, attr)
	}
}

