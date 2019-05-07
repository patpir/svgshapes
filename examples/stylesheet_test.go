package examples

import (
	"log"
	"os"
	"encoding/xml"

	"github.com/patpir/svgshapes"
)

type stylesheet struct {
	// XMLName ensures the creation of a <style> tag for stylesheet objects
	XMLName xml.Name `xml:"style"`
	// Type will always be "text/css" for CSS stylesheets, must be included in the <style> tag
	Type    string   `xml:"type,attr"`
	// Content of style tags in SVG files must be wrapped in CDATA tags
	Content string   `xml:",cdata"`
}

/*
This example shows how to use the document's Shape property to inject
other XML tags, in this example a stylesheet.
It is the responsibility of the injected object (e.g. the stylesheet struct)
to provide XML marshal information.
*/
func ExampleStylesheet() {
	// style defines the content of the <style> tag
	style := `.my-shape { fill: red; }`

	doc := svgshapes.NewDocument("-50,-50,100,100")
	s := &stylesheet{
		Type: "text/css",
		Content: style,
	}
	// Simply append arbitrary objects to the Shapes property.
	// Everything in this slice will be written as a direct child tag of the <svg> tag.
	doc.Shapes = append(doc.Shapes, s)
	// add a shape to the document, to demonstrate that the stylesheet is working
	circle := doc.Circle(10.5, 25, 15)
	circle.Class = "my-shape"

	err := doc.WriteIndent(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	// Output: <svg viewBox="-50,-50,100,100" xmlns="http://www.w3.org/2000/svg">
	//     <style type="text/css"><![CDATA[.my-shape { fill: red; }]]></style>
	//     <circle class="my-shape" cx="10.5" cy="25" r="15"></circle>
	// </svg>
}

