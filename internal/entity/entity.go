package entity

import (
	"errors"
	"math"
	"math/rand"
)

var VertexNotFound = errors.New("err happened during vertex retrieval")

type ConfigDataModel struct {
	Name string
	X, Y float32
}

type Canvas struct {
	Width, Height int
}

const (
	NONE = iota
	CITY
)

type Position struct {
	X, Y float32
}

type Point struct {
	Type int
	Name string
	Position
}

type Graph struct {
	vertices map[Point][]Point
}

func (g *Graph) AddEdge(src, dst Point) {
	if !g.IsEdgeExist(src, dst) {
		g.vertices[src] = append(g.vertices[src], dst)
	}
	if !g.IsEdgeExist(dst, src) {
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

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[Point][]Point),
	}
}

type Iteration struct {
	Path []Point
}

func (it *Iteration) GetFitness() float64 {
	var distance float64
	for i := 0; i < len(it.Path); i++ {
		src := it.Path[i]
		var dst Point
		if i+1 < len(it.Path) {
			dst = it.Path[i+1]
		} else {
			dst = it.Path[0]
		}

		distanceX := float64(src.X - dst.X)
		distanceY := float64(src.Y - dst.Y)

		if distanceX < 0 {
			distanceX = -distanceX
		}
		if distanceY < 0 {
			distanceY = -distanceY
		}

		distance += math.Sqrt((distanceX * distanceX) + (distanceY * distanceY))
	}

	return 1 / distance
}

type Training struct {
	Iterations []*Iteration
}

func (t *Training) GetFittest() *Iteration {
	result := t.Iterations[0]

	for _, iteration := range t.Iterations {
		if iteration.GetFitness() > result.GetFitness() {
			result = iteration
		}
	}

	return result
}

func NewTrainingWithGeneration(src []*ConfigDataModel, generations int) *Training {
	result := new(Training)

	for i := 0; i < generations; i++ {
		iteration := new(Iteration)

		for _, v := range src {
			iteration.Path = append(iteration.Path, Point{
				Name: v.Name,
				Position: Position{
					X: v.X,
					Y: v.Y,
				},
			})
		}

		rand.Shuffle(len(iteration.Path), func(i, j int) { iteration.Path[i], iteration.Path[j] = iteration.Path[j], iteration.Path[i] })

		result.Iterations = append(result.Iterations, iteration)
	}

	return result
}

func NewTraining() *Training {
	return new(Training)
}
