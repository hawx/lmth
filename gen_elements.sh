#!/usr/bin/env bash

# see https://html.spec.whatwg.org/multipage/#toc-semantics
DOCUMENT='html head title style';
SECTIONS='body article section nav aside h1 h2 h3 h4 h5 h6 hgroup header footer address';
GROUPING_CONTENT='p pre blockquote ol ul menu li dl dt dd figure figcaption main search div';
TEXT='a em strong small s cite q dfn abbr ruby rt rp data time code var samp kbd sub sup i b u mark bdi bdo span';
EDITS='ins del';
EMBEDDED_CONTENT='picture iframe object video audio map';
TABLES='table caption colgroup tbody thead tfoot tr td th';
FORMS='form label button select datalist optgroup option textarea output progress meter fieldset legend';
INTERACTIVE='details summary dialog';
SCRIPT='script noscript template slot canvas';

ALL="$DOCUMENT $SECTIONS $GROUPING_CONTENT $TEXT $EDITS $EMBEDDED_CONTENT $TABLES $FORMS $INTERACTIVE $SCRIPT";

cat <<EOF > elements/elements.go
package elements

import "hawx.me/code/lmth"
EOF

for EL in $ALL
do
    NAME="$(tr '[:lower:]' '[:upper:]' <<< ${EL:0:1})${EL:1}"
    cat <<EOF >> elements/elements.go

func $NAME(attr lmth.Attr, children ...lmth.Node) lmth.Node {
  return lmth.Element("$EL", attr, children...)
}
EOF
done

VOID_ELEMENTS='area base br col embed hr img input link meta param source track wbr';

for EL in $VOID_ELEMENTS
do
    NAME="$(tr '[:lower:]' '[:upper:]' <<< ${EL:0:1})${EL:1}"
    cat <<EOF >> elements/elements.go

func $NAME(attr lmth.Attr) lmth.Node {
  return lmth.Element("$EL", attr)
}
EOF
done

go fmt ./elements
