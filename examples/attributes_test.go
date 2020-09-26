package examples

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/patpir/svgshapes"
)

type Scale struct {
	X, Y float64
}

func (s Scale) MarshalXMLAttr(_ xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name: s.XMLName(),
		Value: fmt.Sprintf("scale(%s %s)",
			strconv.FormatFloat(s.X, 'f', -1, 64),
			strconv.FormatFloat(s.Y, 'f', -1, 64),
		),
	}, nil
}

func (s Scale) XMLName() xml.Name {
	return xml.Name{
		Local: "transform",
	}
}

func ExampleAttributes() {
	tf := Scale{
		X: 1.5,
		Y: 1.25,
	}

	doc := svgshapes.NewDocument("-100,-50,200,100")
	// create group with custom transform attribute, and add a circle
	g := doc.Group()
	g.Attributes = append(g.Attributes, tf)
	g.Circle(10.5, 25, 15)
	// create semi-transparent, reference circle without transform
	circle := doc.Circle(10.5, 25, 15)
	circle.SetAttribute("opacity", "0.3")

	err := doc.WriteIndent(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	// Output: <svg viewBox="-100,-50,200,100" xmlns="http://www.w3.org/2000/svg">
	//     <g transform="scale(1.5 1.25)">
	//         <circle cx="10.5" cy="25" r="15"></circle>
	//     </g>
	//     <circle opacity="0.3" cx="10.5" cy="25" r="15"></circle>
	// </svg>
}

