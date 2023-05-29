package graph

import (
	"fmt"
	"sort"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/entity"
	"university/generic_algorithm_project/internal/genetic"
	"university/generic_algorithm_project/internal/tools"

	"github.com/barkimedes/go-deepcopy"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GetGraphSeries() *entity.GetGraphSeriesResponse {
	result := new(entity.GetGraphSeriesResponse)

	geneticAlgorithm := genetic.NewGeneticAlgorithm()

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

	for i := 0; i < config.GetGenerations(); i++ {
		training = geneticAlgorithm.Train(training)

		if bestTraining.GetFittest().GetFitness() < training.GetFittest().GetFitness() {
			bestTraining = deepcopy.MustAnything(training).(*entity.Training)
		}

		result.HistoryRecords = append(result.HistoryRecords, entity.DistanceHistoryRecord{
			Distance:   bestTraining.GetWithLowestDistance().GetDistance(),
			Population: i * training.GetIterationSize(),
		})
	}

	bestIteration := bestTraining.GetFittest()
	result.Fitness = bestIteration.GetFitness()

	for _, v := range bestIteration.Path {
		result.GraphNodes = append(result.GraphNodes, opts.GraphNode{
			Name: v.Name,
			X:    v.X,
			Y:    v.Y,
		})
	}

	for i := 0; i < len(bestIteration.Path)-1; i++ {
		result.GraphLinks = append(result.GraphLinks, opts.GraphLink{
			Source: bestIteration.Path[i].Name,
			Target: bestIteration.Path[i+1].Name,
		})
	}

	sort.Slice(result.HistoryRecords, func(i, j int) bool {
		return result.HistoryRecords[i].Population < result.HistoryRecords[j].Population
	})

	return result
}

func GetLineSeries(src []entity.DistanceHistoryRecord) []opts.LineData {
	var result []opts.LineData

	for _, v := range src {
		result = append(result, opts.LineData{
			Value: []float64{float64(v.Population), v.Distance},
		})
	}

	return result
}

func GetGaugeSeries(src float64) []opts.GaugeData {
	return []opts.GaugeData{{
		Name:  "Best fitness",
		Value: fmt.Sprintf("%.3f", src*100000),
	}}
}
