package dot

import (
	"fmt"
	"strings"
)

type Node struct {
	Name string
	*Attrs
}

func NewNode(name string, attrs ...string) *Node {
	return &Node{
		Name: name,
		Attrs: &Attrs{
			Items: strings.Split(strings.Join(attrs, " "), " "),
		},
	}
}

func (node *Node) AddAttrs(attrs string) {
	node.AddAttrs(attrs)
}

func (node *Node) String() string {
	return fmt.Sprintf("\"%s\" [ %s ]", node.Name, node.Attrs)
}
