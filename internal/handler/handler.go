package handler

import (
	"fmt"
	"net/http"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/tools"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func GetResult(w http.ResponseWriter, r *http.Request) {
	graph := charts.NewGraph()

	canvas := tools.GetCanvas()

	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "TSP",
			Subtitle: "TSP",
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

	var graphNode []opts.GraphNode

	data := tools.GetRandomData(config.GetRandomNames(), canvas.Width, canvas.Height)
	// data := config.GetData()

	for _, v := range data {
		graphNode = append(graphNode, opts.GraphNode{
			Name: v.Name,
			X:    v.X,
			Y:    v.Y,
		})
	}

	graphNode = append(graphNode, opts.GraphNode{
		Name:   "pop",
		X:      250,
		Y:      250,
		Symbol: "none",
	})

	graph.AddSeries("Cities", graphNode, []opts.GraphLink{
		{
			Source: "Kyiv",
			Target: "Zaporizhzhya",
		},
		{
			Source: "Kyiv",
			Target: "pop",
		},
	})

	// // Put data into instance
	// line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
	// 	AddSeries("Category A", generateLineItems()).
	// 	AddSeries("Category B", generateLineItems()).
	// 	SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	graph.Render(w)
}
