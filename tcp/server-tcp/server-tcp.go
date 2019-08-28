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
	//servidor tcp, rodando na porta 8888
	server, err := net.Listen("tcp", ":8888")
	utils.ErrorCheck(err)
	fmt.Println("Server running")

	connection, err := server.Accept()
	utils.ErrorCheck(err)

	fileName, err := bufio.NewReader(connection).ReadString('\n')
	utils.ErrorCheck(err)

	f := strings.Split(fileName, "\n")
	path := fmt.Sprintf("./assets/%s.jpg", f[0])
	file, err := os.Open(path)
	utils.ErrorCheck(err)

	fileBytes, err := ioutil.ReadAll(file)
	utils.ErrorCheck(err)

	_, err = connection.Write(fileBytes)
	utils.ErrorCheck(err)
}
