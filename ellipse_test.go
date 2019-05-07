package svgshapes

import (
	"bytes"
	"testing"
	"encoding/xml"

	"github.com/stretchr/testify/assert"
)

func TestEllipseUsage(t *testing.T) {
	ellipse := Ellipse{
		CenterX: 12.5,
		CenterY: 25.0,
		RadiusX: 37.5,
		RadiusY: 50.0,
	}

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(ellipse)
	assert.Nil(t, err)
	assert.Equal(t, `<ellipse cx="12.5" cy="25" rx="37.5" ry="50"></ellipse>`, buf.String())
}


