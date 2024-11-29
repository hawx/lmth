# lmth

Function based HTML templating for Go. Clearly inspired by Elm, etc.

```go
fruits := []string{"apple", "pear"}

doc := lmth.Body(lmth.Attr{}, 
  lmth.H1(lmth.Attr{"class": "heading-xl"},
    lmth.Text("My web page"),
  ),
  lmth.P(lmth.Attr{},
    lmth.Text("This is some fruit"),
  ),
  lmth.Ul(lmth.Attr{},
    lmth.Map(func(s string) lmth.Node {
      return lmth.Li(lmth.Attr{}, lmth.Text(s))
    }, fruits),
  ),
)

doc.WriteTo(os.Stdout)
```

This is very TODO-level at the moment. Outstanding things are:

- [ ] self-closing tags
- [ ] escaping?
- [ ] a better way to model attributes?
