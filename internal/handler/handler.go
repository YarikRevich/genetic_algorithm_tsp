package handler

import (
	"net/http"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/core"
	"university/generic_algorithm_project/internal/entity"
	"university/generic_algorithm_project/internal/service"
	"university/generic_algorithm_project/internal/tools"

	"github.com/go-echarts/go-echarts/v2/opts"
)

func GetResult(w http.ResponseWriter, r *http.Request) {
	graphRenderer := service.GetGraphRenderer(tools.GetCanvas())

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

	var graphNode []opts.GraphNode

	iteration := training.GetFittest()
	for _, v := range iteration.Path {
		graphNode = append(graphNode, opts.GraphNode{
			Name: v.Name,
			X:    v.X,
			Y:    v.Y,
		})
	}

	var graphLink []opts.GraphLink
	for i := 0; i < len(iteration.Path); i++ {
		var target string
		if i+1 < len(iteration.Path) {
			target = iteration.Path[i+1].Name
		} else {
			target = iteration.Path[0].Name
		}

		graphLink = append(graphLink, opts.GraphLink{
			Source: iteration.Path[0].Name,
			Target: target,
		})
	}

	graphRenderer.AddSeries("Cities", graphNode, graphLink)

	graphRenderer.Render(w)
}
