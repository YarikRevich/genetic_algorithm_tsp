package service

import (
	"fmt"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/core"
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
			Subtitle: fmt.Sprintf("Width: %dpx\n\nHeight: %dpx\n\nCrossover probability: %.2f\n\nCrossover type: %s\n\nMutation probability: %.2f\n\nMutation type: %s\n\nGenerations: %d\n\nAuthor: Yaroslav Svitlytskyi", canvas.Width, canvas.Height, config.GetCrossoverProbability(), config.GetCrossoverType(), config.GetMutationProbability(), config.GetMutationType(), config.GetGenerations()),
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger:   "none",
			TriggerOn: "none",
			Enterable: true,
			Show:      true,
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:    true,
			Feature: &opts.ToolBoxFeature{},
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  fmt.Sprintf("%dpx", canvas.Width),
			Height: fmt.Sprintf("%dpx", canvas.Height),
			Theme:  types.ThemeMacarons,
		}))

	return graph
}

func GetGraphSeries() ([]opts.GraphNode, []opts.GraphLink) {
	geneticAlgorithm := core.NewGeneticAlgorithm()

	generations := config.GetGenerations()

	training := entity.NewTrainingWithGeneration(
		tools.GetData(
			config.GetData(),
			config.GetRandomNames(),
			config.IsRandom(),
		), generations)

	for i := 0; i < generations; i++ {
		training = geneticAlgorithm.Train(training)
	}

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
	for i := 0; i < len(iteration.Path)-1; i++ {
		graphLinks = append(graphLinks, opts.GraphLink{
			Source: iteration.Path[i].Name,
			Target: iteration.Path[i+1].Name,
		})
	}

	return graphNodes, graphLinks
}
