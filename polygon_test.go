package svgshapes

import (
	"bytes"
	"testing"
	"encoding/xml"

	"github.com/stretchr/testify/assert"
)

func TestPolygonUsage(t *testing.T) {
	poly := Polygon{
		Points: []Point{
			Point{  50,   0 },
			Point{ 100, 100 },
		},
	}
	poly.Points = append(poly.Points, Point{   0, 100 })

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(poly)
	assert.Nil(t, err)
	assert.Equal(t, `<polygon points="50,0 100,100 0,100"></polygon>`, buf.String())
}

