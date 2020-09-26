package svgshapes

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Note: this wrapper is actually not required to test the functionality of Shapebase.
However, Shapebase is only used within shape structs just as the one below and it makes more sense to test it in the same way.
*/
type shape struct {
	XMLName xml.Name `xml:"shape"`
	Shapebase
}

func checkShapeXML(t *testing.T, expectedXML string, shape interface{}) {
	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(shape)

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, buf.String())
}

func TestShapeId(t *testing.T) {
	s := shape{}
	s.Id = "shape-identifier"

	checkShapeXML(t, `<shape id="shape-identifier"></shape>`, s)
}

func TestShapeClass(t *testing.T) {
	s := shape{}
	s.Class = "some-class some-other-class"

	checkShapeXML(t, `<shape class="some-class some-other-class"></shape>`, s)
}

func TestShapeStyle(t *testing.T) {
	s := shape{}
	s.Style = "fill: #123456; stroke: hsl(10deg, 10%, 10%)"

	checkShapeXML(t, `<shape style="fill: #123456; stroke: hsl(10deg, 10%, 10%)"></shape>`, s)
}

func TestShapeFill(t *testing.T) {
	s := shape{}
	s.Fill = "red"

	checkShapeXML(t, `<shape fill="red"></shape>`, s)
}

func TestShapeStroke(t *testing.T) {
	s := shape{}
	s.Stroke = "red"

	checkShapeXML(t, `<shape stroke="red"></shape>`, s)
}

func TestShapeSetAttribute(t *testing.T) {
	s := shape{}
	s.SetAttribute("opacity", "0.5")
	s.SetAttribute("opacity", "0.75")

	checkShapeXML(t, `<shape opacity="0.75"></shape>`, s)
}

