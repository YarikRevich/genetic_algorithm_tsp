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
	graphRenderer := service.GetGraphRenderer(tools.GetCanvas())

	graphNodes, graphLinks, fitness := service.GetGraphSeries()
	graphRenderer.AddSeries("Cities", graphNodes, graphLinks, charts.WithGraphChartOpts(
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

	page := components.NewPage()

	gaugeRenderer := service.GetGaugeRenderer()
	gaugeRenderer.SetGlobalOptions(charts.WithGridOpts(opts.Grid{
		Left: "50%",
	}))
	gaugeRenderer.AddSeries("Training result", service.GetGaugeSeries(fitness))

	page.AddCharts(gaugeRenderer)

	barRenderer := service.GetBarRenderer(tools.GetCanvas())
	barRenderer.SetGlobalOptions(charts.WithGridOpts(opts.Grid{
		Right: "50%",
	}))
	page.AddCharts(barRenderer)

	err = page.Render(w)
	if err != nil {
		log.Fatalln(err)
	}
}
