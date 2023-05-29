package genetic

import (
	"math/rand"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/entity"
)

type GeneticAlgorithm struct{}

func (ga *GeneticAlgorithm) getCrossoverOX(src, dst *entity.Iteration) *entity.Iteration {
	result := &entity.Iteration{
		Path: src.Path[:],
	}

	unusedPositions := make(map[int]bool)

	breakingPosition := rand.Intn(len(src.Path))
	for i := 0; i < len(src.Path); i++ {
		var exist bool
		for _, v := range result.Path {
			if v == dst.Path[i] {
				exist = true
			}
		}

		if !exist {
			for unusedPositions[breakingPosition] {
				breakingPosition = (breakingPosition + 1) % len(src.Path)
			}
			result.Path[breakingPosition] = dst.Path[i]
			unusedPositions[breakingPosition] = true
			breakingPosition = (breakingPosition + 1) % len(src.Path)
		}
	}

	return result
}

func (ga *GeneticAlgorithm) getCrossoverCX(src, dst *entity.Iteration) *entity.Iteration {
	result := &entity.Iteration{
		Path: src.Path[:],
	}

	unusedPositions := make(map[int]bool)

	startPoint := rand.Intn(len(src.Path))

main:
	for i := startPoint; i != startPoint; {
		result.Path[i] = src.Path[i]
		unusedPositions[i] = true

		for q, v := range src.Path {
			if v == dst.Path[i] {
				i = q
				continue main
			}
		}

		i = -1
	}

	for i := 0; i < len(src.Path); i++ {
		if !unusedPositions[i] {
			result.Path[i] = dst.Path[i]
		}
	}

	return result
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

func (ga *GeneticAlgorithm) getMutationInversion(iteration *entity.Iteration) *entity.Iteration {
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

func (ga *GeneticAlgorithm) getMutationTransposition(iteration *entity.Iteration) *entity.Iteration {
	result := &entity.Iteration{
		Path: iteration.Path[:],
	}

	src := rand.Intn(len(iteration.Path))
	dst := rand.Intn(len(iteration.Path))

	for src == dst {
		dst = rand.Intn(len(iteration.Path))
	}

	result.Path[src], result.Path[dst] = result.Path[dst], iteration.Path[src]

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

		var mutationResult *entity.Iteration

		switch config.GetMutationType() {
		case config.MUTATION_INVERSION:
			mutationResult = ga.getMutationInversion(crossoverResult)
		case config.MUTATION_TRANSPOSITION:
			mutationResult = ga.getMutationTransposition(crossoverResult)
		}

		result.Iterations = append(result.Iterations, mutationResult)
	}

	return result
}

func NewGeneticAlgorithm() *GeneticAlgorithm {
	return new(GeneticAlgorithm)
}
