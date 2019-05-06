package main

import (
	"log"
	"os"

	"github.com/patpir/svgshapes"
)

func main() {
	circle := svgshapes.Circle{
		CenterX: 10.5,
		CenterY: 25,
		Radius:  15,
	}
	polygon := svgshapes.Polygon{
		Points: []svgshapes.Point{
			svgshapes.Point{ 10, 20 },
			svgshapes.Point{ 4.5, 4.5 },
			svgshapes.Point{ 20, 8 },
		},
	}

	doc := svgshapes.NewDocument("-50,-50,100,100")
	doc.Shapes = append(doc.Shapes, circle)
	doc.Shapes = append(doc.Shapes, polygon)

	err := doc.Write(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

