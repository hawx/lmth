# lmth ([pkg.go.dev](https://pkg.go.dev/hawx.me/code/lmth))

Function based HTML templating for Go. Clearly inspired by Elm, etc.

```go
fruits := []string{"apple", "pear"}

doc := elements.Body(lmth.Attr{},
  elements.H1(lmth.Attr{"class": "heading-xl"},
    lmth.Text("My web page"),
  ),
  elements.P(lmth.Attr{},
    lmth.Text("This is some fruit"),
  ),
  elements.Ul(lmth.Attr{},
    lmth.Map(func(s string) lmth.Node {
      return elements.Li(lmth.Attr{}, lmth.Text(s))
    }, fruits),
  ),
)

doc.WriteTo(os.Stdout)
```

This is very TODO-level at the moment.
