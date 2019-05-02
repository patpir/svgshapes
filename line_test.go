package svgshapes

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineUsage(t *testing.T) {
	line := Line{
		StartX: 0.5,
		StartY: 15,
		EndX:   30,
		EndY:   45,
	}

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(line)
	assert.Nil(t, err)
	assert.Equal(t, `<line x1="0.5" y1="15" x2="30" y2="45"></line>`, buf.String())
}

