package main

import (
	"fmt"
	"network"
	"quorum"
)

func establishQuorum() {
	var port int
	print("Port number: ")
	fmt.Scanf("%d", &port)
	networkServer, err := network.NewTCPServer(port)
	if err != nil {
		println(err)
		return
	}
	s, err := quorum.CreateState(networkServer)
	s.JoinSia()
	select {}
}

func main() {
	var input string
	for {
		print("Please enter a command: ")
		fmt.Scanln(&input)

		switch input {
		default:
			println("unrecognized command")
		case "e":
			establishQuorum()
		case "q":
			return
		}
	}
}
