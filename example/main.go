package main

import (
	"fmt"

	"github.com/zjj2wry/go-dot"
)

func main() {
	g := dot.NewGraph()
	g.AddAttrs(`label="netstat" labeljust="l" fontname="Arial" fontsize="14" rankdir="LR" bgcolor="lightgray" style="solid" penwidth="0.5" pad="0.0" nodesep="0.35"`)
	g.AddGlobalNodeAttrs(`shape="ellipse" style="filled" fillcolor="lightblue" fontname="Verdana" penwidth="1.0" margin="0.05,0.0"`)
	g.AddGlobalEdgeAttrs(`minlen="2" fontsize=12 color="saddlebrown"`)

	subGraph := dot.NewSubGraph()
	subGraph.AddAttrs(`bgcolor="#e6ecfa" label="pod-b" labelloc="t" labeljust="c" fontsize="18"`)
	nodes := []string{"pod-a1", "pod-a2", "pod-a3", "pod-a4", "pod-a5", "pod-b", "pod-c1", "coredns", "mysql", "redis"}
	for _, name := range nodes {
		n := dot.NewNode(name, fmt.Sprintf(`label="%s"`, name))
		subGraph.AddNode(n)
	}
	diamondNodes := []string{"inbound-tcp", "inbound-udp", "outbound-tcp", "outbound-udp"}
	for _, name := range diamondNodes {
		n := dot.NewNode(name, fmt.Sprintf(`shape="diamond" fillcolor="yellow" label="%s" penwidth="0.5"`, name))
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
