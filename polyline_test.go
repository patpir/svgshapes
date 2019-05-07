package svgshapes

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolylineUsage(t *testing.T) {
	poly := Polyline{
		Points: []Point{
			Point{  50,   0 },
			Point{ 100, 100 },
		},
	}
	poly.Points = append(poly.Points, Point{    0, 100 })

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(poly)
	assert.Nil(t, err)
	assert.Equal(t, `<polyline points="50,0 100,100 0,100"></polyline>`, buf.String())
}

