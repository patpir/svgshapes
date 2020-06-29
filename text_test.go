package svgshapes

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextUsage(t *testing.T) {
	text := Text{
		StartX: 12.3,
		StartY: 45,
		Text: "Hello, World!",
	}

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(text)
	assert.Nil(t, err)
	assert.Equal(t, `<text x="12.3" y="45">Hello, World!</text>`, buf.String())
}
