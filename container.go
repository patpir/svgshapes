package svgshapes

import ()

type Container interface {
	Group() *Group
	Circle(cx, cy, radius float64) *Circle
	Polygon(points ...Point) *Polygon
	Line(startX, startY, endX, endY float64) *Line
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

