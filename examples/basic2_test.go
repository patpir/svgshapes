package examples

import (
	"log"
	"os"

	"github.com/patpir/svgshapes"
)

func ExampleBasic2() {
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

	err := doc.WriteIndent(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	// Output: <svg viewBox="-50,-50,100,100" xmlns="http://www.w3.org/2000/svg">
	//     <circle cx="10.5" cy="25" r="15"></circle>
	//     <polygon points="10,20 4.5,4.5 20,8"></polygon>
	// </svg>
}

