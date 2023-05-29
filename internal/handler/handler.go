package handler

import (
	"fmt"
	"log"
	"net/http"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/graph"
	"university/generic_algorithm_project/internal/tools"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func GetResult(w http.ResponseWriter, r *http.Request) {
	canvas := tools.GetCanvas()

	graphRenderer := charts.NewGraph()

	graphRenderer.SetGlobalOptions(
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

	getGraphSeriesResponse := graph.GetGraphSeries()
	graphRenderer.AddSeries(
		"Cities",
		getGraphSeriesResponse.GraphNodes,
		getGraphSeriesResponse.GraphLinks,
		charts.WithGraphChartOpts(
			opts.GraphChart{
				Layout: "none",
				Roam:   true,
				EdgeLabel: &opts.EdgeLabel{
					Show: true,
				},
			}))

	err := graphRenderer.Render(tools.GetOutputWriter())
	if err != nil {
		log.Fatalln(err)
	}

	err = graphRenderer.Render(w)
	if err != nil {
		log.Fatalln(err)
	}

	gaugeRenderer := charts.NewGauge()
	gaugeRenderer.AddSeries(
		"Training result",
		graph.GetGaugeSeries(getGraphSeriesResponse.Fitness))

	err = gaugeRenderer.Render(w)
	if err != nil {
		log.Fatalln(err)
	}

	lineRenderer := charts.NewLine()
	lineRenderer.AddSeries(
		"Distance history",
		graph.GetLineSeries(getGraphSeriesResponse.HistoryRecords),
		charts.WithLineChartOpts(
			opts.LineChart{
				Step: "start",
			}))

	err = lineRenderer.Render(w)
	if err != nil {
		log.Fatalln(err)
	}
}
