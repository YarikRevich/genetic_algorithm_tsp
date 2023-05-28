package graph

import (
	"math"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/entity"
	"university/generic_algorithm_project/internal/tools"
)

var generatedGraph *entity.Graph = entity.NewGraph()

func Init() {
	data := tools.GetData(
		config.GetData(),
		config.GetRandomNames(),
		config.IsRandom(),
	)

	canvas := tools.GetCanvas()

	for h := 0; h < canvas.Height; h++ {
		for w := 0; w < canvas.Width; w++ {
			src := entity.Point{
				Position: entity.Position{
					X: float32(w),
					Y: float32(h),
				},
			}

			for _, v := range data {
				if int(v.X) == w && int(v.Y) == h {
					src.Type = entity.CITY
					break
				}
			}

			if !math.Signbit(float64(h - 1)) {
				top := entity.Point{
					Position: entity.Position{
						X: float32(w),
						Y: float32(h - 1),
					},
				}

				for _, v := range data {
					if int(v.X) == w && int(v.Y) == h-1 {
						top.Type = entity.CITY
						break
					}
				}

				generatedGraph.AddEdge(src, top)
			}

			if !math.Signbit(float64(w - 1)) {
				left := entity.Point{
					Position: entity.Position{
						X: float32(w - 1),
						Y: float32(h),
					},
				}

				for _, v := range data {
					if int(v.X) == w-1 && int(v.Y) == h {
						left.Type = entity.CITY
						break
					}
				}

				generatedGraph.AddEdge(src, left)
			}

			if h+1 < canvas.Height {
				bottom := entity.Point{
					Position: entity.Position{
						X: float32(w),
						Y: float32(h + 1),
					},
				}

				for _, v := range data {
					if int(v.X) == w && int(v.Y) == h+1 {
						bottom.Type = entity.CITY
						break
					}
				}

				generatedGraph.AddEdge(src, bottom)
			}

			if w+1 < canvas.Width {
				right := entity.Point{
					Position: entity.Position{
						X: float32(w + 1),
						Y: float32(h),
					},
				}

				for _, v := range data {
					if int(v.X) == w+1 && int(v.Y) == h {
						right.Type = entity.CITY
						break
					}
				}

				generatedGraph.AddEdge(src, right)
			}
		}
	}
}

func GetGeneratedGraph() *entity.Graph {
	return generatedGraph
}
