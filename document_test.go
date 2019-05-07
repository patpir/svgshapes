package svgshapes

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteEmptyDocument(t *testing.T) {
	doc := NewDocument("0,10,100,200")

	var buf bytes.Buffer
	doc.Write(&buf)
	assert.Equal(t, `<svg viewBox="0,10,100,200" xmlns="http://www.w3.org/2000/svg"></svg>`, buf.String())
}

func TestWriteDocumentWithBaseAttribute(t *testing.T) {
	doc := NewDocument("0,10,100,200")
	doc.Fill = "#123456"

	var buf bytes.Buffer
	doc.Write(&buf)
	assert.Contains(t, buf.String(), `fill="#123456"`)
}

func TestWriteComplexDocument(t *testing.T) {
	doc := NewDocument("0,10,100,200")
	doc.Group()
	doc.Circle(15, 25, 35)
	doc.Polygon(
		Point{ 10, 20 },
		Point{ 30, 40 },
		Point{ 50, 60 },
	)

	var buf bytes.Buffer
	doc.Write(&buf)
	docText := buf.String()
	assert.Contains(t, docText, `viewBox="0,10,100,200"`)
	assert.Contains(t, docText, `xmlns="http://www.w3.org/2000/svg"`)
	assert.Contains(t, docText, "<g></g>")
	assert.Contains(t, docText, "<circle ")
	assert.Contains(t, docText, "<polygon ")
}

