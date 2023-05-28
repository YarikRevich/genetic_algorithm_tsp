package handler

import (
	"log"
	"net/http"
	"university/generic_algorithm_project/internal/service"
	"university/generic_algorithm_project/internal/tools"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GetResult(w http.ResponseWriter, r *http.Request) {
	graphRenderer := service.GetGraphRenderer(tools.GetCanvas())

	graphNodes, graphLinks := service.GetGraphSeries()
	graphRenderer.AddSeries("Cities", graphNodes, graphLinks, charts.WithGraphChartOpts(opts.GraphChart{
		Layout: "none",
		Roam:   true,
		EdgeLabel: &opts.EdgeLabel{
			Show: true,
		},
	}))

	graphRenderer.AddJSFuncs(`function handleButtonClick() { alert('Custom button clicked!'); }`)

	graphRenderer.AddCustomizedJSAssets(`var button = document.createElement("button");
	button.innerHTML = "Custom Button";
	button.addEventListener("click", handleButtonClick);
	document.body.appendChild(button);`)

	err := graphRenderer.Render(w)
	if err != nil {
		log.Fatalln(err)
	}
}
