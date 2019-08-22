package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/PedroCosta8/sistemas-distribuidos/utils"
)

func main() {
	//servidor tcp, rodando na porta 8888
	server, err := net.Listen("tcp", ":8888")
	utils.ErrorCheck(err)
	fmt.Println("Server running")

	defer server.Close()

	TcpSendFile(server)
}

func TcpSendFile(server net.Listener) error {

	connection, err := server.Accept()
	if err != nil {
		return err
	}

	file, err := os.Open("input")
	if err != nil {
		return err
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	//escreve os dados na conexao estabelecida
	_, err = connection.Write(fileBytes)
	if err != nil {
		return err
	}
	return nil
}
