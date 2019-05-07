package svgshapes

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathUsageAbsolute(t *testing.T) {
	path := Path{}
	path.MoveTo(12.5, 25.0)
	path.LineTo(37.5, 50.0)
	path.HLineTo(62.5)
	path.VLineTo(75.0)
	path.CubicBezierTo(12.5, 25.0, 37.5, 50.0, 62.5, 75.0)
	path.SmoothCubicBezierTo(12.5, 25.0, 37.5, 50.0)
	path.QuadraticBezierTo(12.5, 25.0, 37.5, 50.0)
	path.SmoothQuadraticBezierTo(62.5, 75.0)
	path.EllipticalArcTo(12.5, 25.0, 37.5, true, false, 50.0, 62.5)
	path.ClosePath()

	d := strings.TrimSpace(strings.Replace(`
M 12.5 25
L 37.5 50
H 62.5
V 75
C 12.5 25 37.5 50 62.5 75
S 12.5 25 37.5 50
Q 12.5 25 37.5 50
T 62.5 75
A 12.5 25 37.5 1 0 50 62.5
Z
	`, "\n", " ", -1))

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(path)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<path d="%s"></path>`, d), buf.String())
}

func TestPathUsageRelative(t *testing.T) {
	path := Path{}
	path.MoveRelative(12.5, 25.0)
	path.LineRelative(37.5, 50.0)
	path.HLineRelative(62.5)
	path.VLineRelative(75.0)
	path.CubicBezierRelative(12.5, 25.0, 37.5, 50.0, 62.5, 75.0)
	path.SmoothCubicBezierRelative(12.5, 25.0, 37.5, 50.0)
	path.QuadraticBezierRelative(12.5, 25.0, 37.5, 50.0)
	path.SmoothQuadraticBezierRelative(62.5, 75.0)
	path.EllipticalArcRelative(12.5, 25.0, 37.5, true, false, 50.0, 62.5)
	path.ClosePath()

	d := strings.TrimSpace(strings.Replace(`
m 12.5 25
l 37.5 50
h 62.5
v 75
c 12.5 25 37.5 50 62.5 75
s 12.5 25 37.5 50
q 12.5 25 37.5 50
t 62.5 75
a 12.5 25 37.5 1 0 50 62.5
Z
	`, "\n", " ", -1))

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	err := enc.Encode(path)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<path d="%s"></path>`, d), buf.String())
}

