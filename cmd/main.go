package main

import (
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/graph"
	"university/generic_algorithm_project/internal/server"
	"university/generic_algorithm_project/internal/tools"
)

func init() {
	config.Init()
	graph.Init()
}

func main() {
	server.Run()

	tools.WaitForExit()
}
