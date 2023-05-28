package main

import (
	"math/rand"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/graph"
	"university/generic_algorithm_project/internal/server"
	"university/generic_algorithm_project/internal/tools"
)

func init() {
	config.Init()
	graph.Init()

	rand.Seed(tools.GetRandSeed())
}

func main() {
	server.Run()

	tools.WaitForExit()
}
