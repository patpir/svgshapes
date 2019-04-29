SVG Shapes
==========

Golang library to define and encode shapes as SVG.


Examples
--------

You can easily define individual shapes like this:

```go
circle := svgshapes.Circle{
	CenterX: 10.5,
	CenterY: 25,
	Radius:  15,
}
polygon := svgshapes.Polygon{
	points: []svgshapes.Point{
		svgshapes.Point{ 10, 20 },
		svgshapes.Point{ 4.5, 4.5 },
		svgshapes.Point{ 20, 8 },
	},
}
```

Add these shapes to a document, which can then be encoded as an SVG file:

```go
doc := svgshapes.NewDocument("-50,-50,50,50")
doc.Shapes = append(doc.Shapes, circle)
doc.Shapes = append(doc.Shapes, polygon)

err := doc.Write(file)
if err != nil {
	log.Fatal(err)
}
```

The document also provides methods to create shapes directly.
The following example produces the same result as before:

```go
doc := svgshapes.NewDocument("-50,-50,50,50")
doc.Circle(10.5, 25, 15)
doc.Polygon(
	svgshapes.Point{ 10, 20 },
	svgshapes.Point{ 4.5, 4.5 },
	svgshapes.Point{ 20, 8 },
)

err := doc.Write(file)
if err != nil {
	log.Fatal(err)
}
```

