package svgshapes

import (
	"bytes"
	"fmt"
	"strconv"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) String() string {
	return fmt.Sprintf("%s,%s",
		strconv.FormatFloat(p.X, 'f', -1, 64),
		strconv.FormatFloat(p.Y, 'f', -1, 64),
	)
}

type PointList []Point

func (points PointList) MarshalText() (text []byte, err error) {
	var pointText [][]byte
	for _, point := range points {
		pointText = append(pointText, []byte(point.String()))
	}

	return bytes.Join(pointText, []byte(" ")), nil
}

