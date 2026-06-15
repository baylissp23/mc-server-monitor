package main

import (
	"serverStatusMonitor/internal/monitor"
)

var resultsChannel = make(chan monitor.Server)

func main() {
	servers := []string{
		"bee.mc-complex.com:25565",
		"buzz.lemoncloud.org:25565",
		"buzz.mysticmc.co:25565"}

	for i := range servers {
		go monitor.PingServer(servers[i], resultsChannel)
	}

	serversResults := []monitor.Server{}
	for range servers {
		receivedServer := <-resultsChannel
		serversResults = append(serversResults, receivedServer)
	}

	monitor.PrintItems(serversResults)
}
