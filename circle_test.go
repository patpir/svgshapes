package svgshapes

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleUsage(t *testing.T) {
	circle := Circle{
		CenterX: 12.3,
		CenterY: 45,
		Radius:  6.7,
	}

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(circle)
	assert.Nil(t, err)
	assert.Equal(t, `<circle cx="12.3" cy="45" r="6.7"></circle>`, buf.String())
}

