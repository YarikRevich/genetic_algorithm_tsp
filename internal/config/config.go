package config

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"university/generic_algorithm_project/internal/entity"

	"github.com/0chain/common/core/viper"
)

var (
	ConfigFieldNotFound  = errors.New("err happened during config field retrieval")
	ConfigFieldWrongType = errors.New("err happened during config field conversion")
)

var (
	configFile = flag.String("configFile", "./config/cities.yaml", "Describes a config file used for TSP algorithm")
	random     = flag.Bool("random", false, "Enables random data generation")
	output     = flag.String("output", "output.html", "Describes the path of the output")
)

var (
	generations          int
	representation       string
	elitism              bool
	crossoverProbability float64
	crossoverType        string
	mutationProbability  float64
	mutationType         string

	data []*entity.ConfigDataModel

	randomCanvas entity.Canvas
	randomNames  []string
)

const RESULT_PATH = "/"

func Init() {
	flag.Parse()

	err := viper.ReadConfigFile(*configFile)
	if err != nil {
		log.Fatalln(err)
	}

	generations = viper.GetInt("meta.generations")
	representation = viper.GetString("meta.representation")
	elitism = viper.GetBool("meta.elitism")
	crossoverProbability = viper.GetFloat64("meta.crossover.probability")
	crossoverType = viper.GetString("meta.crossover.type")
	mutationProbability = viper.GetFloat64("meta.mutation.probability")
	mutationType = viper.GetString("meta.mutation.type")

	dataRaw, ok := viper.Get("data").([]interface{})
	if !ok {
		log.Fatalln(ConfigFieldNotFound.Error())
	}

	for i := range dataRaw {
		name := viper.GetString(fmt.Sprintf("data.%d.name", i))
		x := viper.GetFloat64(fmt.Sprintf("data.%d.x", i))
		y := viper.GetFloat64(fmt.Sprintf("data.%d.y", i))

		data = append(data, &entity.ConfigDataModel{
			Name: name,
			X:    float32(x),
			Y:    float32(y),
		})
	}

	randomCanvasRaw, ok := viper.Get("random.canvas").(map[string]interface{})
	if !ok {
		log.Fatalln(ConfigFieldNotFound.Error())
	}

	widthRaw, ok := randomCanvasRaw["width"]
	if !ok {
		log.Fatalln(ConfigFieldNotFound.Error())
	}

	widthInt, ok := widthRaw.(int)
	if !ok {
		log.Fatalln(ConfigFieldWrongType.Error())
	}

	randomCanvas.Width = widthInt

	heightRaw, ok := randomCanvasRaw["height"]
	if !ok {
		log.Fatalln(ConfigFieldNotFound.Error())
	}

	heightInt, ok := heightRaw.(int)
	if !ok {
		log.Fatalln(ConfigFieldWrongType.Error())
	}

	randomCanvas.Height = heightInt

	randomNames = viper.GetStringSlice("random.names")
}

func GetOutput() string {
	return *output
}

func IsRandom() bool {
	return *random
}

func GetGenerations() int {
	return generations
}

func GetRepresentation() string {
	return representation
}

func IsElitism() bool {
	return elitism
}

func GetCrossoverProbability() float64 {
	return crossoverProbability
}

func GetCrossoverType() string {
	return crossoverType
}

func GetMutationProbability() float64 {
	return mutationProbability
}

func GetMutationType() string {
	return mutationType
}

func GetData() []*entity.ConfigDataModel {
	return data
}

func GetRandomCanvas() entity.Canvas {
	return randomCanvas
}

func GetRandomNames() []string {
	return randomNames
}
