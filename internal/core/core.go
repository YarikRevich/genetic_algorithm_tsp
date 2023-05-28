package core

import (
	"fmt"
	"math/rand"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/entity"
)

type GeneticAlgorithm struct {
}

func (ga *GeneticAlgorithm) getCrossover(src, dst *entity.Iteration) *entity.Iteration {
	result := &entity.Iteration{
		Path: src.Path,
	}

	// Child Tour
	// c := base.Tour{}
	// c.InitTour(size)

	// Number of crossover
	nc := int(config.GetCrossoverProbability() * float64(len(src.Path)))
	// Start positions of cross over for parent 1
	sp := int(rand.Float32() * float32(len(src.Path)))
	// End position of cross over for parent 1
	ep := (sp + nc) % len(src.Path)
	// Parent 2 slots
	p2s := make([]int, 0, len(src.Path)-nc)
	// Populate child with parent 1
	if sp < ep {
		for i := 0; i < len(src.Path); i++ {
			if i >= sp && i < ep {
				result.Path = append(result.Path, src.Path[i])
			} else {
				p2s = append(p2s, i)
			}
		}
	} else if sp > ep {
		for i := 0; i < len(src.Path); i++ {
			if !(i >= ep && i < sp) {
				result.Path = append(result.Path, src.Path[i])
			} else {
				p2s = append(p2s, i)
			}
		}
	}

	j := 0
	for i := 0; i < len(src.Path); i++ {
		var exist bool

		for _, v := range result.Path {
			if v == dst.Path[i] {
				exist = true
			}
		}
		if !exist {
			result.Path[p2s[j]] = dst.Path[i]
			j++
		}
	}
	return result
}

func (ga *GeneticAlgorithm) getMutation(iteration *entity.Iteration) *entity.Iteration {
	result := &entity.Iteration{
		Path: iteration.Path[:],
	}

	for src := range iteration.Path {
		if rand.Float64() < config.GetMutationProbability() {
			dst := int(float64(len(iteration.Path)) * rand.Float64())

			result.Path[src], result.Path[dst] = result.Path[dst], result.Path[src]
		}
	}

	return result
}

func (ga *GeneticAlgorithm) getTournamentSelection(src *entity.Training) *entity.Iteration {
	trainingTemp := entity.NewTraining()

	for i := 0; i < len(src.Iterations); i++ {
		r := int(rand.Float64() * float64(len(src.Iterations)))
		trainingTemp.Iterations = append(trainingTemp.Iterations, src.Iterations[r])
	}

	return trainingTemp.GetFittest()
}

func (ga *GeneticAlgorithm) Train(src *entity.Training) *entity.Training {
	result := entity.NewTraining()

	if config.IsElitism() {
		result.Iterations = append(result.Iterations, src.GetFittest())
	}

	fmt.Println(ga.getTournamentSelection(src), ga.getTournamentSelection(src))

	for i := 0; i < len(src.Iterations)-1; i++ {

		mutation := ga.getMutation(ga.getCrossover(
			ga.getTournamentSelection(src),
			ga.getTournamentSelection(src)))
		result.Iterations = append(result.Iterations, mutation)
	}

	return result
}

func NewGeneticAlgorithm() *GeneticAlgorithm {
	return new(GeneticAlgorithm)
}
