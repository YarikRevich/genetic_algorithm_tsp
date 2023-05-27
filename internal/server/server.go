package server

import (
	"log"
	"net"
	"net/http"
	"university/generic_algorithm_project/internal/config"
	"university/generic_algorithm_project/internal/handler"
	readinessprobe "university/generic_algorithm_project/internal/readiness_probe"
	"university/generic_algorithm_project/internal/tools"

	"github.com/pkg/browser"
)

func Run() {
	http.HandleFunc(config.RESULT_PATH, handler.GetResult)

	listener, err := net.Listen("tcp", "localhost:")
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		err = http.Serve(listener, nil)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	err = readinessprobe.Run(listener.Addr().String())
	if err != nil {
		log.Fatalln(err)
	}

	err = browser.OpenURL(tools.GetLocalServerURL(listener.Addr().String()))
	if err != nil {
		log.Fatalln(err)
	}
}
