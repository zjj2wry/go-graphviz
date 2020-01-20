package dot

import "strings"

type Node struct {
	Name  string
	*Attrs
}

func NewNode(name string, attrs ...string) Node {
	return Node{
		Name: name,
		Attrs: &Attrs{
			Items: strings.Split(strings.Join(attrs, " "), " "),
		},
	}
}

func (node *Node)AddAttrs(attrs string){
	node.AddAttrs(attrs)
}

func(node *Node)String()string{
	return fmt.Sprintf("\"%s\" -> \"%s\" [%s]", node.Src, node.Dst, node.Attrs)
}