package handler

import (
	"net/http"
	"university/generic_algorithm_project/internal/service"
	"university/generic_algorithm_project/internal/tools"
)

func GetResult(w http.ResponseWriter, r *http.Request) {
	graphRenderer := service.GetGraphRenderer(tools.GetCanvas())

	graphNodes, graphLinks := service.GetGraphSeries()
	graphRenderer.AddSeries("Cities", graphNodes, graphLinks)

	graphRenderer.Render(w)
}
