package tools

import (
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/entity"
)

var randSeed = time.Now().Unix()

func GetRandSeed() int64 {
	return randSeed
}

func GetLocalServerURL(address string) string {
	var url url.URL
	url.Scheme = "http"
	url.Host = address

	return url.String()
}

func GetCanvas() entity.Canvas {
	if config.IsRandom() {
		return config.GetRandomCanvas()
	}

	var result entity.Canvas

	data := config.GetData()
	for _, v := range data {
		if v.X > float32(result.Width) {
			result.Width = int(v.X)
		}

		if v.Y > float32(result.Height) {
			result.Height = int(v.Y)
		}
	}

	result.Width *= 2
	result.Height *= 2

	return result
}

func GetData(data []*entity.ConfigDataModel, randomData []string, isRandom bool) []*entity.ConfigDataModel {
	var result []*entity.ConfigDataModel

	canvas := GetCanvas()

	if isRandom {
		usedPositions := make(map[entity.Position]bool)
		for _, v := range randomData {
			ticker := time.NewTicker(time.Millisecond * 20)
			for range ticker.C {
				ticker.Stop()

				generatedPosition := entity.Position{
					X: float32(rand.Intn(canvas.Width)),
					Y: float32(rand.Intn(canvas.Height)),
				}

				if _, ok := usedPositions[generatedPosition]; ok {
					ticker.Reset(time.Millisecond * 20)
					continue
				}

				result = append(result, &entity.ConfigDataModel{
					Name: v,
					X:    generatedPosition.X,
					Y:    generatedPosition.Y,
				})

				break
			}
		}
	} else {
		result = data
	}

	return result
}

func WaitForExit() {
	log.Println("Press 'Ctrl+C' to stop")

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(time.Millisecond * 500)
	for range ticker.C {
		select {
		case <-exitCh:
			os.Exit(0)
		default:
		}
	}
}
