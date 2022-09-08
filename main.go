package main

import (
	"backup/client"
	"backup/server"
	"os"
	"strconv"
)

func main() {
	num := len(os.Args)
	if num < 3 {
		println("Usage Server: backup server <port>")
		println("Usage Client: backup client <upload_folder> <data_folder>")
		return
	}
	if os.Args[1] == "server" {
		port, err := strconv.Atoi(os.Args[2])
		if err != nil {
			println("Invalid port")
			return
		}
		server.GetServer(port)
		return
	}
	if os.Args[1] != "client" {
		println("Invalid type")
		return
	}
	if num < 4 {
		println("Usage Client: backup client <upload_folder> <data_folder>")
		return
	}
	client := client.GetClient(os.Args[2], os.Args[3])
	client.Start()
}
