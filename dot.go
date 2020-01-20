package main

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

type Node struct {
	Name  string
	Attrs *Attrs
}

type Edge struct {
	Src   string
	Dst   string
	Attrs *Attrs
}

type Attrs struct {
	Iterms []string
}

func (o *Attrs) String() string {
	return strings.Join(o.Iterms, " ")
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

func NewNode(name string, attrs ...string) Node {
	return Node{
		Name: name,
		Attrs: &Attrs{
			Iterms: strings.Split(strings.Join(attrs, " "), " "),
		},
	}
}

func NewEdge(src, dst string, attrs ...string) Edge {
	return Edge{
		Src: src,
		Dst: dst,
		Attrs: &Attrs{
			Iterms: strings.Split(strings.Join(attrs, " "), " "),
		},
	}
}

func (g *Graph) AddAttrs(attrs string) {
	g.Attrs.Iterms = append(g.Attrs.Iterms, strings.Split(attrs, " ")...)
}

func (g *Graph) AddGlobalNodeAttrs(attrs string) {
	g.NodeAttrs.Iterms = append(g.NodeAttrs.Iterms, strings.Split(attrs, " ")...)
}

func (g *Graph) AddGlobalEdgeAttrs(attrs string) {
	g.EdgeAttrs.Iterms = append(g.EdgeAttrs.Iterms, strings.Split(attrs, " ")...)
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

func (o *Graph) String() string {
	result := fmt.Sprintf("%s %s{\n", o.Type, o.Name)
	if o.Attrs != nil {
		result += fmt.Sprintf("\t%s\n", o.Attrs)
	}

	if o.NodeAttrs != nil {
		if len(o.NodeAttrs.Iterms) > 0 {
			result += fmt.Sprintf("\tnode [%s]\n", o.NodeAttrs)
		}
	}

	if o.EdgeAttrs != nil {
		if len(o.NodeAttrs.Iterms) > 0 {
			result += fmt.Sprintf("\tedge [%s]\n", o.EdgeAttrs)
		}
	}

	for _, node := range o.Nodes {
		result += fmt.Sprintf("\t\"%s\" [%s]\n", node.Name, node.Attrs)
	}

	for _, subgraph := range o.Subgraphs {
		items := strings.Split(subgraph.String(), "\n")
		for _, item := range items {
			result += fmt.Sprintf("\t%s\n", item)
		}
	}

	for _, edge := range o.Edges {
		result += fmt.Sprintf("\t\"%s\" -> \"%s\" [%s]\n", edge.Src, edge.Dst, edge.Attrs)
	}

	result += fmt.Sprintf("}\n")

	return result
}

func main() {
	g := NewGraph()
	g.AddAttrs(`label="netstat" labeljust="l" fontname="Arial" fontsize="14" rankdir="LR" bgcolor="lightgray" style="solid" penwidth="0.5" pad="0.0" nodesep="0.35"`)
	g.AddGlobalNodeAttrs(`shape="ellipse" style="filled" fillcolor="lightblue" fontname="Verdana" penwidth="1.0" margin="0.05,0.0"`)
	g.AddGlobalEdgeAttrs(`minlen="2" fontsize=12 color="saddlebrown"`)

	subGraph := NewSubGraph()
	subGraph.AddAttrs(`bgcolor="#e6ecfa" label="pod-b" labelloc="t" labeljust="c" fontsize="18"`)
	nodes := []string{"pod-a1", "pod-a2", "pod-a3", "pod-a4", "pod-a5", "pod-b", "pod-c1", "coredns", "mysql", "redis"}
	for _, name := range nodes {
		n := NewNode(name, fmt.Sprintf(`label="%s"`, name))
		subGraph.AddNode(n)
	}
	diamondNodes := []string{"inbound-tcp", "inbound-udp", "outbound-tcp", "outbound-udp"}
	for _, name := range diamondNodes {
		n := NewNode(name, fmt.Sprintf(`shape="diamond" fillcolor="yellow" label="%s" penwidth="0.5"`, name))
		subGraph.AddNode(n)
	}
	g.AddSubgraph(subGraph)

	g.AddEdge(NewEdge("pod-a1", "inbound-tcp", `label="ESTABLISH=50\nTIME_WAIT=10000"`))
	g.AddEdge(NewEdge("pod-a2", "inbound-tcp", `label="ESTABLISH=50\nTIME_WAIT=10000"`))
	g.AddEdge(NewEdge("pod-a3", "inbound-tcp", `label="ESTABLISH=50\nTIME_WAIT=10000"`))
	g.AddEdge(NewEdge("pod-a4", "inbound-udp", `label="ESTABLISH=50\nTIME_WAIT=10000"`))
	g.AddEdge(NewEdge("pod-a5", "inbound-udp", `label="ESTABLISH=50\nTIME_WAIT=10000"`))
	g.AddEdge(NewEdge("inbound-tcp", "pod-b", `label="ESTABLISH=50\nTIME_WAIT=10000"`))
	g.AddEdge(NewEdge("inbound-udp", "pod-b", `label="ESTABLISH=50\nTIME_WAIT=10000"`))
	g.AddEdge(NewEdge("pod-b", "outbound-udp", `label="ESTABLISH=50\nTIME_WAIT=10000"`))
	g.AddEdge(NewEdge("pod-b", "outbound-tcp", `label="ESTABLISH=50\nTIME_WAIT=10000"`))
	g.AddEdge(NewEdge("outbound-tcp", "mysql", `label="ESTABLISH=50\nTIME_WAIT=10000"`))
	g.AddEdge(NewEdge("outbound-tcp", "redis", `label="ESTABLISH=50\nTIME_WAIT=10000"`))
	g.AddEdge(NewEdge("outbound-tcp", "pod-c1", `label="ESTABLISH=50\nTIME_WAIT=10000"`))
	g.AddEdge(NewEdge("outbound-udp", "coredns", `label="ESTABLISH=50\nTIME_WAIT=10000"`))

	fmt.Println(g.String())
}
