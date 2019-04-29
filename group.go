package svgshapes

import (
	"encoding/xml"
)

type Group struct {
	XMLName xml.Name `xml:"g"`
	Shapebase
	container
}

func (group *Group) Group() *Group {
	subGroup := &Group{}
	group.Shapes = append(group.Shapes, subGroup)
	return subGroup
}

func (group *Group) Circle(cx, cy, radius float64) *Circle {
	circle := &Circle{
		CenterX: cx,
		CenterY: cy,
		Radius:  radius,
	}
	group.Shapes = append(group.Shapes, circle)
	return circle
}

func (group *Group) Polygon(points ...Point) *Polygon {
	p := &Polygon{
		Points: points,
	}
	group.Shapes = append(group.Shapes, p)
	return p
}

