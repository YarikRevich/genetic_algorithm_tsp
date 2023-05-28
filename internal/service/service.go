package service

import (
	"fmt"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/entity"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func GetGraphRenderer(canvas entity.Canvas) *charts.Graph {
	graph := charts.NewGraph()

	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "TSP",
			Subtitle: fmt.Sprintf("Width: %dpx\n\nHeight: %dpx\n\nCrossover probability: %.2f\n\nCrossover type: %s\n\nMutation probability: %.2f\n\nMutation type: %s\n\nAuthori: Yaroslav Svitlytskyi", canvas.Width, canvas.Height, config.GetCrossoverProbability(), config.GetCrossoverType(), config.GetMutationProbability(), config.GetMutationType()),
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger:   "item",
			TriggerOn: "click",
			Enterable: true,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  fmt.Sprintf("%dpx", canvas.Width),
			Height: fmt.Sprintf("%dpx", canvas.Height),
			Theme:  types.ThemeMacarons,
		}))

	return graph
}

func GetPathToPoint() {

}

// # Python3 program to implement traveling salesman
// # problem using naive approach.
// from sys import maxsize
// from itertools import permutations
// V = 4

// # implementation of traveling Salesman Problem
// def travellingSalesmanProblem(graph, s):

// 	# store all vertex apart from source vertex
// 	vertex = []
// 	for i in range(V):
// 		if i != s:
// 			vertex.append(i)

// 	# store minimum weight Hamiltonian Cycle
// 	min_path = maxsize
// 	next_permutation=permutations(vertex)
// 	for i in next_permutation:

// 		# store current Path weight(cost)
// 		current_pathweight = 0

// 		# compute current path weight
// 		k = s
// 		for j in i:
// 			current_pathweight += graph[k][j]
// 			k = j
// 		current_pathweight += graph[k][s]

// 		# update minimum
// 		min_path = min(min_path, current_pathweight)

// 	return min_path

// # Driver Code
// if __name__ == "__main__":

// 	# matrix representation of graph
// 	graph = [[0, 10, 15, 20], [10, 0, 35, 25],
// 			[15, 35, 0, 30], [20, 25, 30, 0]]
// 	s = 0
// 	print(travellingSalesmanProblem(graph, s))
