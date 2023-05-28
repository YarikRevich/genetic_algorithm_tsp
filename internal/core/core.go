package core

import (
	"math/rand"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/entity"
)

type GeneticAlgorithm struct{}

func (ga *GeneticAlgorithm) getCrossoverOX(src, dst *entity.Iteration) *entity.Iteration {
	return nil
}

func (ga *GeneticAlgorithm) getCrossoverCX(src, dst *entity.Iteration) *entity.Iteration {
	return nil
}

func (ga *GeneticAlgorithm) getCrossoverPBC(src, dst *entity.Iteration) *entity.Iteration {
	result := &entity.Iteration{
		Path: src.Path[:],
	}

	nc := int(config.GetCrossoverProbability() * float64(len(src.Path)))
	sp := int(rand.Float32() * float32(len(src.Path)))
	ep := (sp + nc) % len(src.Path)
	selected := make([]int, 0, len(src.Path)-nc)
	if sp < ep {
		for i := 0; i < len(src.Path); i++ {
			if !(i >= sp && i < ep) {
				selected = append(selected, i)
			}
		}
	} else if sp > ep {
		for i := 0; i < len(src.Path); i++ {
			if i >= ep && i < sp {
				selected = append(selected, i)
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
			result.Path[selected[j]] = dst.Path[i]
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

	for i := 0; i < len(src.Iterations)-1; i++ {
		var crossoverResult *entity.Iteration

		switch config.GetCrossoverType() {
		case config.CROSSOVER_OX:
			crossoverResult = ga.getCrossoverOX(
				ga.getTournamentSelection(src),
				ga.getTournamentSelection(src))
		case config.CROSSOVER_CX:
			crossoverResult = ga.getCrossoverCX(
				ga.getTournamentSelection(src),
				ga.getTournamentSelection(src))
		case config.CROSSOVER_PBC:
			crossoverResult = ga.getCrossoverPBC(
				ga.getTournamentSelection(src),
				ga.getTournamentSelection(src))
		}

		result.Iterations = append(result.Iterations, ga.getMutation(crossoverResult))
	}

	return result
}

func NewGeneticAlgorithm() *GeneticAlgorithm {
	return new(GeneticAlgorithm)
}
