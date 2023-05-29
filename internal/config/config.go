package config

import (
	"flag"
	"fmt"
	"log"
	"time"
	"university/generic_algorithm_project/internal/entity"

	"github.com/0chain/common/core/viper"
	"github.com/pkg/errors"
)

var (
	ConfigFieldNotFound       = errors.New("err happened during config field retrieval")
	ConfigFieldWrongType      = errors.New("err happened during config field conversion")
	ConfigWrongParameterValue = errors.New("err happened during parameter value check")
)

var (
	configFile = flag.String("configFile", "./config/config.yaml", "Describes a config file used for TSP algorithm")
	random     = flag.Bool("random", false, "Enables random data generation")
	cities     = flag.Int("cities", 10, "Describes the number of cities used to generate when random data generation is enabled")
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
	showHistory          bool
	historyDelay         time.Duration

	data []*entity.ConfigDataModel

	randomCanvas entity.Canvas
	randomNames  []string
)

const (
	CROSSOVER_OX  = "OX"
	CROSSOVER_CX  = "CX"
	CROSSOVER_PBC = "PBC"
)

const (
	MUTATION_INVERSION     = "inversion"
	MUTATION_TRANSPOSITION = "transposition"
)

const RESULT_PATH = "/result"

func Init() {
	flag.Parse()

	err := viper.ReadConfigFile(*configFile)
	if err != nil {
		log.Fatalln(err)
	}

	generations = viper.GetInt("meta.generations")
	if generations == 0 {
		log.Fatalln(errors.Wrap(ConfigWrongParameterValue, "meta.generations"))
	}

	representation = viper.GetString("meta.representation")
	elitism = viper.GetBool("meta.elitism")
	crossoverProbability = viper.GetFloat64("meta.crossover.probability")
	if crossoverProbability == 0 {
		log.Fatalln(errors.Wrap(ConfigWrongParameterValue, "meta.crossover.probability"))
	}

	crossoverType = viper.GetString("meta.crossover.type")
	if crossoverType != CROSSOVER_OX && crossoverType != CROSSOVER_CX && crossoverType != CROSSOVER_PBC {
		log.Fatalln(errors.Wrap(ConfigWrongParameterValue, "meta.crossover.type"))
	}

	mutationProbability = viper.GetFloat64("meta.mutation.probability")
	if mutationProbability == 0 {
		log.Fatalln(errors.Wrap(ConfigWrongParameterValue, "meta.mutation.probability"))
	}

	mutationType = viper.GetString("meta.mutation.type")
	if mutationType != MUTATION_INVERSION && mutationType != MUTATION_TRANSPOSITION {
		log.Fatalln(errors.Wrap(ConfigWrongParameterValue, "meta.mutation.type"))
	}

	showHistory = viper.GetBool("meta.view.history.show")
	historyDelay = viper.GetDuration("meta.view.history.delay")
	if showHistory && historyDelay == 0 {
		log.Fatalln(errors.Wrap(ConfigWrongParameterValue, "meta.view.history.delay"))
	}

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
	if *cities > len(randomNames) {
		log.Fatalln(errors.Wrap(ConfigWrongParameterValue, "cities number is bigger than the number of available cities"))
	}
}

func GetOutput() string {
	return *output
}

func IsRandom() bool {
	return *random
}

func GetCities() int {
	return *cities
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

func IsShowHistory() bool {
	return showHistory
}

func GetHistoryDelay() time.Duration {
	return historyDelay
}
