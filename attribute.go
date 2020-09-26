package svgshapes

import (
	"encoding/xml"
)

type Attribute interface {
	xml.MarshalerAttr
	XMLName() xml.Name
}

type rawAttribute xml.Attr

var _ Attribute = rawAttribute{}

func (a rawAttribute) MarshalXMLAttr(_ xml.Name) (xml.Attr, error) {
	return xml.Attr(a), nil
}

func (a rawAttribute) XMLName() xml.Name {
	return a.Name
}

