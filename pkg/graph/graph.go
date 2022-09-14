package graph

import (
	"fmt"
	"strings"
)

type NodeId = int

type Graph interface {
	GetNode(id NodeId) Node
	GetNodes() []Node
	GetArcsFrom(id NodeId) []*Arc
	NodeCount() int
	ArcCount() int
	AsString() string
	EstimateDistance(source, target NodeId) int
	SetArcFlags(nodeId NodeId, flag bool)
	EnableAllArcs()
}

type DynamicGraph interface {
	Graph
	AddNode(n Node)
	AddEdge(edge Edge) bool
	AddArc(from, to NodeId, distance int) bool
}

type Node struct {
	Lon float64
	Lat float64
	// TODO: id?
	// TODO: Point attribute/ implement Point type
}

type Edge struct {
	From     NodeId
	To       NodeId
	Distance int
	arcFlag  bool
}

type Arc struct {
	To       NodeId
	Distance int
	arcFlag  bool
}

type Arcs = []Arc

func NewNode(lon float64, lat float64) *Node {
	return &Node{Lon: lon, Lat: lat}
}

func NewEdge(to, from NodeId, distance int, arcFlag bool) *Edge {
	return &Edge{To: to, From: from, Distance: distance, arcFlag: arcFlag}
}
func MakeEdge(from, to NodeId, distance int, arcFlag bool) Edge {
	return Edge{To: to, From: from, Distance: distance, arcFlag: arcFlag}
}

func NewArc(to NodeId, distance int, arcFlag bool) *Arc {
	return &Arc{To: to, Distance: distance, arcFlag: arcFlag}
}

func MakeArc(to NodeId, distance int, arcFlag bool) Arc {
	return Arc{To: to, Distance: distance, arcFlag: arcFlag}
}

func (e Edge) Destination() NodeId {
	return e.To
}

func (e Edge) Cost() int {
	return e.Distance
}

func (e Edge) Invert() Edge {
	return Edge{From: e.To, To: e.From, Distance: e.Distance}
}

func (e Edge) toArc() *Arc {
	return NewArc(e.To, e.Distance, e.ArcFlag())
}

func (e Edge) ArcFlag() bool {
	return e.arcFlag
}

func (e *Edge) SetArcFlag(flag bool) {
	e.arcFlag = flag
}

func (a Arc) toEdge(from NodeId) *Edge {
	return NewEdge(from, a.To, a.Distance, a.ArcFlag())
}

func (a Arc) Destination() NodeId {
	return a.To
}

func (a Arc) Cost() int {
	return a.Distance
}

func (a Arc) ArcFlag() bool { return a.arcFlag }

func (a *Arc) SetArcFlag(flag bool) { a.arcFlag = flag }

func GraphAsString(g Graph) string {
	var sb strings.Builder

	// write number of nodes and number of edges
	sb.WriteString(fmt.Sprintf("%v\n", g.NodeCount()))
	sb.WriteString(fmt.Sprintf("%v\n", g.ArcCount()))

	// list all nodes structured as "id lat lon"
	for i := 0; i < g.NodeCount(); i++ {
		node := g.GetNode(i)
		sb.WriteString(fmt.Sprintf("%v %v %v\n", i, node.Lat, node.Lon))
	}

	// list all edges structured as "fromId targetId distance"
	for i := 0; i < g.NodeCount(); i++ {
		for _, arc := range g.GetArcsFrom(i) {
			sb.WriteString(fmt.Sprintf("%v %v %v\n", i, arc.Destination(), arc.Cost()))
		}
	}
	return sb.String()
}
