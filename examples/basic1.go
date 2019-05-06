package main

import (
	"log"
	"os"

	"github.com/patpir/svgshapes"
)

func main() {
	doc := svgshapes.NewDocument("-50,-50,100,100")
	doc.Circle(10.5, 25, 15)
	doc.Polygon(
		svgshapes.Point{ 10, 20 },
		svgshapes.Point{ 4.5, 4.5 },
		svgshapes.Point{ 20, 8 },
	)

	err := doc.Write(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

