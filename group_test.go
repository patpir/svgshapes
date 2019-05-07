package svgshapes

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupXML(t *testing.T) {
	group := &Group{}

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(group)
	assert.Nil(t, err)
	assert.Equal(t, `<g></g>`, buf.String())
}

func TestGroupWithGroupXML(t *testing.T) {
	group := &Group{}
	group.Group()

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(group)
	assert.Nil(t, err)
	assert.Equal(t, `<g><g></g></g>`, buf.String())
}

func TestGroupWithGroupWithShapeXML(t *testing.T) {
	group := &Group{}
	subGroup := group.Group()
	subGroup.Circle(1.0, 2.0, 3.0)

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(group)
	assert.Nil(t, err)
	assert.Equal(t, `<g><g><circle cx="1" cy="2" r="3"></circle></g></g>`, buf.String())
}

