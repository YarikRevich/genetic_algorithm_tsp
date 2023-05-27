package entity

import "errors"

var VertexNotFound = errors.New("err happened during vertex retrieval")

type ConfigDataModel struct {
	Name string
	X, Y float32
}

type Canvas struct {
	Width, Height int
}

const (
	CITY = iota
	NONE
)

type Position struct {
	X, Y float32
}

type Point struct {
	Type int
	Position
}

type Graph struct {
	vertices map[Point][]Point
}

func (g *Graph) AddEdge(src, dst Point) {
	if g.IsEdgeExist(src, dst) {
		g.vertices[src] = append(g.vertices[src], dst)
	}
	if g.IsEdgeExist(dst, src) {
		g.vertices[dst] = append(g.vertices[dst], src)
	}
}

func (g *Graph) IsEdgeExist(src, dst Point) bool {
	connections, ok := g.vertices[src]
	if !ok {
		return false
	}
	for _, v := range connections {
		if v == dst {
			return false
		}
	}
	return true
}
