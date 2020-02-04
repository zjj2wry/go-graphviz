package dot

import (
	"fmt"
	"strings"
)

type Edge struct {
	Src string
	Dst string
	*Attrs
}

func NewEdge(src, dst string, attrs ...string) *Edge {
	return &Edge{
		Src: src,
		Dst: dst,
		Attrs: &Attrs{
			Items: strings.Split(strings.Join(attrs, " "), " "),
		},
	}
}

func (edge *Edge) AddAttrs(attrs string) {
	edge.AddAttrs(attrs)
}

func (edge *Edge) String() string {
	return fmt.Sprintf("\"%s\" -> \"%s\" [ %s ]", edge.Src, edge.Dst, edge.Attrs)
}
