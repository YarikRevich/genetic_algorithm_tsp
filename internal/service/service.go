package service

import (
	"fmt"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/entity"
	"university/generic_algorithm_project/internal/genetic"
	"university/generic_algorithm_project/internal/history"
	"university/generic_algorithm_project/internal/tools"

	"github.com/barkimedes/go-deepcopy"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func GetGraphRenderer(canvas entity.Canvas) *charts.Graph {
	graph := charts.NewGraph()

	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "TSP",
			Subtitle: fmt.Sprintf("Width: %dpx\n\nHeight: %dpx\n\nCrossover probability: %.2f\n\nCrossover type: %s\n\nMutation probability: %.2f\n\nMutation type: %s\n\nGenerations: %d\n\nRandom seek: %d\n\nAuthor: Yaroslav Svitlytskyi", canvas.Width, canvas.Height, config.GetCrossoverProbability(), config.GetCrossoverType(), config.GetMutationProbability(), config.GetMutationType(), config.GetGenerations(), tools.GetRandSeed()),
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
			PageTitle: "TSP",
			Width:     fmt.Sprintf("%dpx", canvas.Width),
			Height:    fmt.Sprintf("%dpx", canvas.Height),
			Theme:     types.ThemeMacarons,
		}))

	// graph.AddJSFuncs(`
	// function handleButtonClick()
	// { alert('Custom button clicked!');};
	// var button = document.createElement("button");
	// button.innerHTML = "YARIK";
	// button.style.size = "200px";
	// button.addEventListener("click", handleButtonClick);
	// document.body.appendChild(button);`)

	return graph
}

func GetGraphSeries() ([]opts.GraphNode, []opts.GraphLink, float64, *history.DistanceHistoryTracker) {
	geneticAlgorithm := genetic.NewGeneticAlgorithm()

	generations := config.GetGenerations()

	distanceHistoryTracker := history.NewDistanceHistoryTracker()

	bestTraining := entity.NewTrainingWithGeneration(
		tools.GetData(
			config.GetData(),
			config.GetRandomNames(),
			config.IsRandom(),
		))

	training := entity.NewTrainingWithGeneration(
		tools.GetData(
			config.GetData(),
			config.GetRandomNames(),
			config.IsRandom(),
		))

	for i := 0; i < generations; i++ {
		training = geneticAlgorithm.Train(training)

		if bestTraining.GetFittest().GetFitness() < training.GetFittest().GetFitness() {
			bestTraining = deepcopy.MustAnything(training).(*entity.Training)
		}

		distanceHistoryTracker.AddRecord(entity.DistanceHistoryRecord{
			Distance:   training.GetWithLowestDistance().GetDistance(),
			Population: i * training.GetIterationSize(),
		})
	}

	var graphNodes []opts.GraphNode

	iteration := bestTraining.GetFittest()
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

	return graphNodes, graphLinks, iteration.GetFitness(), distanceHistoryTracker
}

func GetLineRenderer(canvas entity.Canvas) *charts.Line {
	return charts.NewLine()
}

func GetLineSeries(distanceHistoryTracker *history.DistanceHistoryTracker) []opts.LineData {
	var result []opts.LineData

	for _, v := range distanceHistoryTracker.GetAllOrdered() {
		result = append(result, opts.LineData{
			Value: []float64{v.Distance, float64(v.Population)},
		})
	}

	return result
}

func GetGaugeRenderer() *charts.Gauge {
	return charts.NewGauge()
}

func GetGaugeSeries(fitness float64) []opts.GaugeData {
	return []opts.GaugeData{{
		Name:  "Best fitness",
		Value: fmt.Sprintf("%.3f", fitness*100000),
	}}
}
