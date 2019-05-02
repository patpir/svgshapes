package svgshapes

import (
	"bytes"
	"testing"
	"encoding/xml"

	"github.com/stretchr/testify/assert"
)

func TestSimpleRectUsage(t *testing.T) {
	rect := Rect{
		StartX: 12.5,
		StartY: 25.0,
		Width:  37.5,
		Height: 50.0,
	}

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(rect)
	assert.Nil(t, err)
	assert.Equal(t, `<rect x="12.5" y="25" width="37.5" height="50"></rect>`, buf.String())
}

func TestRoundedRectUsage(t *testing.T) {
	rect := Rect{
		StartX:  12.5,
		StartY:  25.0,
		Width:   37.5,
		Height:  50.0,
		RadiusX: 10.0,
		RadiusY: 20.0,
	}

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(rect)
	assert.Nil(t, err)
	assert.Equal(t, `<rect x="12.5" y="25" width="37.5" height="50" rx="10" ry="20"></rect>`, buf.String())
}

