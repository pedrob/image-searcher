package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"github.com/PedroCosta8/sistemas-distribuidos/utils"
)

func main() {
	server, err := net.Listen("tcp", ":8888")
	defer server.Close()
	utils.ErrorCheck(err)
	fmt.Println("Server running")

	for {
		connection, err := server.Accept()
		utils.ErrorCheck(err)
		go handleRequest(connection)
	}
}

func handleRequest(connection net.Conn) {
	defer connection.Close()
	fileName, err := bufio.NewReader(connection).ReadString('\n')
	utils.ErrorCheck(err)

	f := strings.Split(fileName, "\n")
	fLow := strings.ToLower(f[0])
	path := fmt.Sprintf("./assets/%s.jpg", fLow)
	file, err := os.Open(path)
	utils.ErrorCheck(err)

	fileBytes, err := ioutil.ReadAll(file)
	utils.ErrorCheck(err)

	_, err = connection.Write(fileBytes)
	utils.ErrorCheck(err)
}
