package service

import (
	"fmt"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/entity"
	"university/generic_algorithm_project/internal/tools"

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

func GetGraphSeries() ([]opts.GraphNode, []opts.GraphLink) {
	// geneticAlgorithm := core.NewGeneticAlgorithm()

	generations := config.GetGenerations()

	training := entity.NewTrainingWithGeneration(
		tools.GetData(
			config.GetData(),
			config.GetRandomNames(),
			config.IsRandom(),
		), generations)

	// fmt.Println(training.Iterations)
	for _, v := range training.Iterations {
		fmt.Println(v.Path)
	}

	// for i := 0; i < generations; i++ {
	// 	training = geneticAlgorithm.Train(training)
	// }

	var graphNodes []opts.GraphNode

	iteration := training.GetFittest()
	for _, v := range iteration.Path {
		graphNodes = append(graphNodes, opts.GraphNode{
			Name: v.Name,
			X:    v.X,
			Y:    v.Y,
		})
	}

	var graphLinks []opts.GraphLink
	for i := 0; i < len(iteration.Path); i++ {
		var target string
		if i+1 < len(iteration.Path) {
			target = iteration.Path[i+1].Name
		} else {
			target = iteration.Path[0].Name
		}

		graphLinks = append(graphLinks, opts.GraphLink{
			Source: iteration.Path[0].Name,
			Target: target,
		})
	}

	return graphNodes, graphLinks
}
