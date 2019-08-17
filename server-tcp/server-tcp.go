package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/PedroCosta8/sistemas-distribuidos/utils"
)

func main() {
	server, err := net.Listen("tcp", ":8888")
	utils.ErrorCheck(err)
	fmt.Println("Server running")

	defer server.Close()

	receiveFileFromClient(server, "./outputFile")
}

func receiveFileFromClient(server net.Listener, filePath string) {
	newFile, err := os.Create(filePath)
	utils.ErrorCheck(err)
	defer newFile.Close()

	connection, err := server.Accept()
	utils.ErrorCheck(err)

	io.Copy(newFile, connection)
	utils.ErrorCheck(err)

	return
}
