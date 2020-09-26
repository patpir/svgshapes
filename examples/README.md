Examples for using svgshapes
============================

To build and run an example use the following procedure:

```shell
go build <example-name>.go
./<example-name>     # on Linux or
<example-name>.exe   # on Windows
```

List of examples:
 - **basic1**: The first example listed in the main README.
               Add shapes using the document's `Circle` and `Polygon` methods.
 - **basic2**: The second example listed in the main README.
               Create shape objects directly and then add them to a document via
               the `Shapes` property.
 - **stylesheet**: Create a custom XML tag for stylesheets and add it to the
                   document's `Shapes` property. Technically, the `Shapes`
                   property holds any sub-elements of a document or group, it is
                   not limited to the shape elements defined in the `svgshapes`
                   package.
 - **attributes**: Create a custom `transform` attribute and add it to a
                   shape group by directly modifying the `Attributes` field.
                   Add an `opacity` attribute via `SetAttribute`.

