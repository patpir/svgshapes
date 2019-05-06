SVG Shapes
==========

Golang library to define and encode shapes as SVG.


Examples
--------

Use the `NewDocument` function to create a document with the given view box.
Shapes can be created and attached to this document with the document's
methods `Circle`, `Polygon` etc.

```go
doc := svgshapes.NewDocument("-50,-50,100,100")
doc.Circle(10.5, 25, 15)
doc.Polygon(
	svgshapes.Point{ 10, 20 },
	svgshapes.Point{ 4.5, 4.5 },
	svgshapes.Point{ 20, 8 },
)
```

To generate an SVG file from the document, use the `Write` method:

```go
err := doc.Write(file)
if err != nil {
	log.Fatal(err)
}
```

You can also define individual shapes without a document:

```go
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
```

Add these shapes to a document by directly modifying the `Shapes` property:

```go
doc := svgshapes.NewDocument("-50,-50,100,100")
doc.Shapes = append(doc.Shapes, circle)
doc.Shapes = append(doc.Shapes, polygon)
```


Limitations
-----------

This SVG library is kept very simple.
Some things might not work as expected:

 - SVG shapes without child elements (e.g. `circle`, `polygon`, `path` etc.)
   are normally written as self-closing tags (tags in the form
   `<circle ... />`, `<polygon ... />`, `<path ... />` etc.).
   Golang's XML encoding library does not support creating self-closing tags
   (see [the corresponding issue](https://github.com/golang/go/issues/21399)
   for details).
 - The SVG specification recommends *"that SVG generators split long path data
   strings across multiple lines, with each line not exceeding 255 characters"*.
   Golang's XML encoding library replaces all newline characters within
   attributes with the entity `&#xA;`.
   As this does not help with viewing or editing such path data,
   **no line-breaks will be inserted** into the path data.
 - Only the most basic SVG elements and their attributes are supported.
 - All numeric values are encoded as unit-less numbers. This means that lengths
   and coordinates will be interpreted in the current user coordinate system.
   Units, such as `px`, `pt`, `em` and `%`, are not supported at the moment.

