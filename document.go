package svgshapes

import (
	"io"
	"encoding/xml"
)

type Document struct {
	XMLName   xml.Name `xml:"svg"`
	Shapebase
	container
	ViewBox   string `xml:"viewBox,attr"`
	Namespace string `xml:"xmlns,attr"`
}

func NewDocument(viewBox string) *Document {
	return &Document{
		ViewBox: viewBox,
		Namespace: "http://www.w3.org/2000/svg",
	}
}

func (doc *Document) Write(writer io.Writer) error {
	enc := xml.NewEncoder(writer)
	return enc.Encode(doc)
}

