package main

import (
	"io"
	"net"
	"os"

	"github.com/PedroCosta8/sistemas-distribuidos/utils"
)

func main() {
	connection, err := net.Dial("tcp", ":8888")
	utils.ErrorCheck(err)
	defer connection.Close()

	sendFile(connection, "./client-tcp/filetxt")
}

func sendFile(connection net.Conn, filePath string) {
	file, err := os.Open(filePath)
	utils.ErrorCheck(err)
	defer file.Close()

	_, err = io.Copy(connection, file)
	utils.ErrorCheck(err)

	return
}
