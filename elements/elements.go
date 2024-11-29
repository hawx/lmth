package elements

import "hawx.me/code/lmth"

func Html(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("html", attr, children...)
}

func Head(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("head", attr, children...)
}

func Title(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("title", attr, children...)
}

func Style(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("style", attr, children...)
}

func Body(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("body", attr, children...)
}

func Article(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("article", attr, children...)
}

func Section(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("section", attr, children...)
}

func Nav(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("nav", attr, children...)
}

func Aside(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("aside", attr, children...)
}

func H1(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("h1", attr, children...)
}

func H2(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("h2", attr, children...)
}

func H3(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("h3", attr, children...)
}

func H4(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("h4", attr, children...)
}

func H5(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("h5", attr, children...)
}

func H6(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("h6", attr, children...)
}

func Hgroup(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("hgroup", attr, children...)
}

func Header(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("header", attr, children...)
}

func Footer(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("footer", attr, children...)
}

func Address(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("address", attr, children...)
}

func P(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("p", attr, children...)
}

func Pre(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("pre", attr, children...)
}

func Blockquote(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("blockquote", attr, children...)
}

func Ol(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("ol", attr, children...)
}

func Ul(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("ul", attr, children...)
}

func Menu(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("menu", attr, children...)
}

func Li(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("li", attr, children...)
}

func Dl(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("dl", attr, children...)
}

func Dt(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("dt", attr, children...)
}

func Dd(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("dd", attr, children...)
}

func Figure(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("figure", attr, children...)
}

func Figcaption(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("figcaption", attr, children...)
}

func Main(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("main", attr, children...)
}

func Search(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("search", attr, children...)
}

func Div(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("div", attr, children...)
}

func A(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("a", attr, children...)
}

func Em(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("em", attr, children...)
}

func Strong(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("strong", attr, children...)
}

func Small(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("small", attr, children...)
}

func S(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("s", attr, children...)
}

func Cite(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("cite", attr, children...)
}

func Q(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("q", attr, children...)
}

func Dfn(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("dfn", attr, children...)
}

func Abbr(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("abbr", attr, children...)
}

func Ruby(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("ruby", attr, children...)
}

func Rt(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("rt", attr, children...)
}

func Rp(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("rp", attr, children...)
}

func Data(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("data", attr, children...)
}

func Time(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("time", attr, children...)
}

func Code(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("code", attr, children...)
}

func Var(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("var", attr, children...)
}

func Samp(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("samp", attr, children...)
}

func Kbd(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("kbd", attr, children...)
}

func Sub(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("sub", attr, children...)
}

func Sup(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("sup", attr, children...)
}

func I(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("i", attr, children...)
}

func B(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("b", attr, children...)
}

func U(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("u", attr, children...)
}

func Mark(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("mark", attr, children...)
}

func Bdi(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("bdi", attr, children...)
}

func Bdo(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("bdo", attr, children...)
}

func Span(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("span", attr, children...)
}

func Ins(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("ins", attr, children...)
}

func Del(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("del", attr, children...)
}

func Picture(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("picture", attr, children...)
}

func Iframe(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("iframe", attr, children...)
}

func Object(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("object", attr, children...)
}

func Video(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("video", attr, children...)
}

func Audio(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("audio", attr, children...)
}

func Map(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("map", attr, children...)
}

func Table(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("table", attr, children...)
}

func Caption(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("caption", attr, children...)
}

func Colgroup(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("colgroup", attr, children...)
}

func Tbody(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("tbody", attr, children...)
}

func Thead(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("thead", attr, children...)
}

func Tfoot(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("tfoot", attr, children...)
}

func Tr(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("tr", attr, children...)
}

func Td(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("td", attr, children...)
}

func Th(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("th", attr, children...)
}

func Form(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("form", attr, children...)
}

func Label(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("label", attr, children...)
}

func Button(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("button", attr, children...)
}

func Select(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("select", attr, children...)
}

func Datalist(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("datalist", attr, children...)
}

func Optgroup(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("optgroup", attr, children...)
}

func Option(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("option", attr, children...)
}

func Textarea(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("textarea", attr, children...)
}

func Output(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("output", attr, children...)
}

func Progress(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("progress", attr, children...)
}

func Meter(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("meter", attr, children...)
}

func Fieldset(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("fieldset", attr, children...)
}

func Legend(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("legend", attr, children...)
}

func Details(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("details", attr, children...)
}

func Summary(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("summary", attr, children...)
}

func Dialog(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("dialog", attr, children...)
}

func Script(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("script", attr, children...)
}

func Noscript(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("noscript", attr, children...)
}

func Template(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("template", attr, children...)
}

func Slot(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("slot", attr, children...)
}

func Canvas(attr lmth.Attr, children ...lmth.Node) lmth.Node {
	return lmth.Element("canvas", attr, children...)
}

func Area(attr lmth.Attr) lmth.Node {
	return lmth.Element("area", attr)
}

func Base(attr lmth.Attr) lmth.Node {
	return lmth.Element("base", attr)
}

func Br(attr lmth.Attr) lmth.Node {
	return lmth.Element("br", attr)
}

func Col(attr lmth.Attr) lmth.Node {
	return lmth.Element("col", attr)
}

func Embed(attr lmth.Attr) lmth.Node {
	return lmth.Element("embed", attr)
}

func Hr(attr lmth.Attr) lmth.Node {
	return lmth.Element("hr", attr)
}

func Img(attr lmth.Attr) lmth.Node {
	return lmth.Element("img", attr)
}

func Input(attr lmth.Attr) lmth.Node {
	return lmth.Element("input", attr)
}

func Link(attr lmth.Attr) lmth.Node {
	return lmth.Element("link", attr)
}

func Meta(attr lmth.Attr) lmth.Node {
	return lmth.Element("meta", attr)
}

func Param(attr lmth.Attr) lmth.Node {
	return lmth.Element("param", attr)
}

func Source(attr lmth.Attr) lmth.Node {
	return lmth.Element("source", attr)
}

func Track(attr lmth.Attr) lmth.Node {
	return lmth.Element("track", attr)
}

func Wbr(attr lmth.Attr) lmth.Node {
	return lmth.Element("wbr", attr)
}
