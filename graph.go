package dot

import (
	"fmt"
	"strings"
)

type Graph struct {
	Type      string
	Name      string
	Attrs     *Attrs
	NodeAttrs *Attrs
	EdgeAttrs *Attrs
	Subgraphs []*Graph
	Nodes     []Node
	Edges     []Edge
}

func NewGraph() *Graph {
	return &Graph{
		Type:      "digraph",
		Name:      "G",
		Attrs:     &Attrs{},
		NodeAttrs: &Attrs{},
		EdgeAttrs: &Attrs{},
		Edges:     []Edge{},
		Subgraphs: []*Graph{},
	}
}

func NewSubGraph() *Graph {
	return &Graph{
		Type:      "subgraph",
		Name:      `"cluster"`,
		Attrs:     &Attrs{},
		NodeAttrs: &Attrs{},
		EdgeAttrs: &Attrs{},
		Edges:     []Edge{},
		Subgraphs: []*Graph{},
	}
}

func (g *Graph) AddAttrs(attrs string) {
	g.Attrs.AddAttrs(attrs)
}

func (g *Graph) AddGlobalNodeAttrs(attrs string) {
	g.NodeAttrs.AddAttrs(attrs)
}

func (g *Graph) AddGlobalEdgeAttrs(attrs string) {
	g.EdgeAttrs.AddAttrs(attrs)
}

func (g *Graph) AddSubgraph(graph *Graph) {
	g.Subgraphs = append(g.Subgraphs, graph)
}

func (g *Graph) AddNode(node Node) {
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) AddEdge(edge Edge) {
	g.Edges = append(g.Edges, edge)
}

func (g *Graph) SetName(name string) {
	g.Name = name
}

func (g *Graph) SetType(graphType string) {
	g.Type = graphType
}

func (o *Graph) String() string {
	result := fmt.Sprintf("%s %s{\n", o.Type, o.Name)
	if o.Attrs != nil {
		result += fmt.Sprintf("\t%s\n", o.Attrs)
	}

	if o.NodeAttrs != nil {
		if len(o.NodeAttrs.Items) > 0 {
			result += fmt.Sprintf("\tnode [%s]\n", o.NodeAttrs)
		}
	}

	if o.EdgeAttrs != nil {
		if len(o.NodeAttrs.Items) > 0 {
			result += fmt.Sprintf("\tedge [%s]\n", o.EdgeAttrs)
		}
	}

	for _, node := range o.Nodes {
		result += fmt.Sprintf("\t%s\n", node)
	}

	for _, subgraph := range o.Subgraphs {
		items := strings.Split(subgraph.String(), "\n")
		for _, item := range items {
			result += fmt.Sprintf("\t%s\n", item)
		}
	}

	for _, edge := range o.Edges {
		result += fmt.Sprintf("\t%s\n", edge)
	}

	result += fmt.Sprintf("}\n")

	return result
}
