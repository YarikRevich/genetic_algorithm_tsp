package handler

import (
	"log"
	"net/http"
	"university/generic_algorithm_project/internal/service"
	"university/generic_algorithm_project/internal/tools"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GetResult(w http.ResponseWriter, r *http.Request) {
	page := components.NewPage()

	// grid.SetGlobalOptions(
	// 	charts.WithGridOpts(opts.Grid{Width: "50%"}),
	// )

	graphRenderer := service.GetGraphRenderer(tools.GetCanvas())

	graphNodes, graphLinks := service.GetGraphSeries()
	graphRenderer.AddSeries("Cities", graphNodes, graphLinks, charts.WithGraphChartOpts(
		opts.GraphChart{
			Layout: "none",
			Roam:   true,
			EdgeLabel: &opts.EdgeLabel{
				Show: true,
			},
		}))

	err := graphRenderer.Render(service.GetOutputWriter())
	if err != nil {
		log.Fatalln(err)
	}

	page.AddCharts(graphRenderer)

	gaugeRenderer := service.GetGaugeRenderer()
	gaugeRenderer.AddSeries("Training result", service.GetGaugeSeries(100))

	page.AddCharts(gaugeRenderer)

	err = page.Render(w)
	if err != nil {
		log.Fatalln(err)
	}
}
