package server

import (
	"backup/util"
	"sync"
)

func GetServer(port int) util.FileServer {
	wg := sync.WaitGroup{}
	wg.Add(1)
	fileserver := util.GetFileServer(port, "fullchain.pem", "privkey.pem")
	wg.Wait()
	return fileserver
}
