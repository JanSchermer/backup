package util

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type FileServer struct {
	port   int
	server http.Server
	auths  Database
}

func GetFileServer(port int, cert string, key string) FileServer {
	server := http.Server{}
	server.Addr = ":" + strconv.Itoa(port)
	if cert != "" && key != "" {
		go server.ListenAndServeTLS(cert, key)
	} else {
		go server.ListenAndServe()
	}
	auths := GetDatabase("auths.txt", ":")
	fileServer := FileServer{port, server, auths}
	server.Handler = http.HandlerFunc(fileServer.Recieve)
	return fileServer
}

func (f *FileServer) Recieve(writer http.ResponseWriter, request *http.Request) {
	file := request.Header.Get("File")
	auth := request.Header.Get("Authorization")
	local := f.auths.Get(auth)
	if local == "" {
		writer.WriteHeader(403)
		return
	}
	file = local + "/" + file
	os.MkdirAll(filepath.Dir(file), 0777)
	fileWriter, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		println(err.Error())
		writer.WriteHeader(500)
		return
	}
	body := request.Body
	io.Copy(fileWriter, body)
	writer.WriteHeader(200)
	writer.Write([]byte("OK"))
}
