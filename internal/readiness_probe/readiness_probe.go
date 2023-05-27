package readinessprobe

import (
	"net"
	"time"
)

func Run(address string) error {
	listener, err := net.DialTimeout("tcp", address, time.Second*5)
	if err != nil {
		return err
	}

	return listener.Close()
}
