package svgshapes

import ()

type Container interface {
	Group() *Group
	Circle(cx, cy, radius float64) *Circle
	Polygon(points ...Point) *Polygon
	Line(startX, startY, endX, endY float64) *Line
	Polyline(points ...Point) *Polyline
	Rect(startX, startY, width, height float64) *Rect
	Ellipse(centerX, centerY, radiusX, radiusY float64) *Ellipse
}

type container struct {
	Shapes []interface{}
}

func (c *container) Group() *Group {
	group := &Group{}
	c.Shapes = append(c.Shapes, group)
	return group
}

func (c *container) Circle(cx, cy, radius float64) *Circle {
	circle := &Circle{
		CenterX: cx,
		CenterY: cy,
		Radius:  radius,
	}
	c.Shapes = append(c.Shapes, circle)
	return circle
}

func (c *container) Polygon(points ...Point) *Polygon {
	polygon := &Polygon{
		Points: points,
	}
	c.Shapes = append(c.Shapes, polygon)
	return polygon
}

func (c *container) Line(startX, startY, endX, endY float64) *Line {
	line := &Line{
		StartX: startX,
		StartY: startY,
		EndX:   endX,
		EndY:   endY,
	}
	c.Shapes = append(c.Shapes, line)
	return line
}

func (c *container) Polyline(points ...Point) *Polyline {
	polyline := &Polyline{
		Points: points,
	}
	c.Shapes = append(c.Shapes, polyline)
	return polyline
}

func (c *container) Rect(startX, startY, width, height float64) *Rect {
	rect := &Rect{
		StartX: startX,
		StartY: startY,
		Width:  width,
		Height: height,
	}
	c.Shapes = append(c.Shapes, rect)
	return rect
}

func (c *container) Ellipse(centerX, centerY, radiusX, radiusY float64) *Ellipse {
	ellipse := &Ellipse{
		CenterX: centerX,
		CenterY: centerY,
		RadiusX: radiusX,
		RadiusY: radiusY,
	}
	c.Shapes = append(c.Shapes, ellipse)
	return ellipse
}

