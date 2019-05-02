package svgshapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerGroup(t *testing.T) {
	c := container{}
	c.Group()

	assert.Equal(t, 1, len(c.Shapes))
	assert.Equal(t, &Group{}, c.Shapes[0])
}

func TestContainerCircle(t *testing.T) {
	c := container{}
	c.Circle(1.2, 3.4, 5.6)

	assert.Equal(t, 1, len(c.Shapes))
	assert.Equal(t, &Circle{
		CenterX: 1.2,
		CenterY: 3.4,
		Radius:  5.6,
	}, c.Shapes[0])
}

func TestContainerPolygon(t *testing.T) {
	c := container{}
	c.Polygon(
		Point{ 1.5, 2.5 },
		Point{ 3.5, 4.5 },
		Point{ 5.5, 6.5 },
	)

	assert.Equal(t, 1, len(c.Shapes))
	assert.Equal(t, &Polygon{
		Points: []Point{
			{ 1.5, 2.5 },
			{ 3.5, 4.5 },
			{ 5.5, 6.5 },
		},
	}, c.Shapes[0])
}

func TestContainerLine(t *testing.T) {
	c := container{}
	c.Line(1.5, 2.5, 3.5, 4.5)

	assert.Equal(t, 1, len(c.Shapes))
	assert.Equal(t, &Line{
		StartX: 1.5,
		StartY: 2.5,
		EndX:   3.5,
		EndY:   4.5,
	}, c.Shapes[0])
}

func TestContainerPolyline(t *testing.T) {
	c := container{}
	c.Polyline(
		Point{ 1.5, 2.5 },
		Point{ 3.5, 4.5 },
		Point{ 5.5, 6.5 },
	)

	assert.Equal(t, 1, len(c.Shapes))
	assert.Equal(t, &Polyline{
		Points: []Point{
			{ 1.5, 2.5 },
			{ 3.5, 4.5 },
			{ 5.5, 6.5 },
		},
	}, c.Shapes[0])
}

func TestContainerRect(t *testing.T) {
	c := container{}
	c.Rect(1.5, 2.5, 3.5, 4.5)

	assert.Equal(t, 1, len(c.Shapes))
	assert.Equal(t, &Rect{
		StartX:  1.5,
		StartY:  2.5,
		Width:   3.5,
		Height:  4.5,
		RadiusX: 0.0,
		RadiusY: 0.0,
	}, c.Shapes[0])
}

func TestContainerEllipse(t *testing.T) {
	c := container{}
	c.Ellipse(1.5, 2.5, 3.5, 4.5)

	assert.Equal(t, 1, len(c.Shapes))
	assert.Equal(t, &Ellipse{
		CenterX: 1.5,
		CenterY: 2.5,
		RadiusX: 3.5,
		RadiusY: 4.5,
	}, c.Shapes[0])
}

