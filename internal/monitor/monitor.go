package monitor

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type Reporter interface {
	Stringer() string
}

type Server struct {
	Address string
	Online  bool
	Latency int64
}

func (server Server) Stringer() string {
	b, _ := json.MarshalIndent(server, "", "  ")
	return string(b)
}

func PrintItems[T Reporter](items []T) {
	for _, item := range items {
		fmt.Println(item.Stringer())
	}
}

func PingServer(address string, ch chan<- Server) {
	var startTime = time.Now()

	conn, err1 := net.DialTimeout("tcp", address, 3*time.Second)
	if err1 != nil {
		serverResult := Server{address, false, 0}
		ch <- serverResult
		return
	}

	var latency = time.Since(startTime)
	err2 := conn.Close()
	if err2 != nil {
		serverResult := Server{address, false, 0}
		ch <- serverResult
		return
	}

	serverResult := Server{address, true, latency.Milliseconds()}
	ch <- serverResult
}
