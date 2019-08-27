package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"github.com/PedroCosta8/sistemas-distribuidos/utils"
)

type Protocol struct {
	WithError bool
	Data      []byte
}

func main() {
	//servidor tcp, rodando na porta 8888
	server, err := net.Listen("tcp", ":8888")
	utils.ErrorCheck(err)
	fmt.Println("Server running")
	//defer server.Close()

	connection, err := server.Accept()
	utils.ErrorCheck(err)

	var protocol Protocol

	fileName, err := bufio.NewReader(connection).ReadString('\n')
	utils.ErrorCheck(err)

	f := strings.Split(fileName, "\n")
	file, err := os.Open(f[0])
	if err != nil {
		protocol.WithError = true
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		protocol.WithError = true
	}

	protocol.Data = fileBytes
	jsonProtocol, err := json.Marshal(protocol)
	utils.ErrorCheck(err)

	response := []byte(jsonProtocol)
	_, err = connection.Write(response)
	utils.ErrorCheck(err)
}
