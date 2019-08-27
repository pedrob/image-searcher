package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/PedroCosta8/sistemas-distribuidos/utils"
)

type Protocol struct {
	WithError bool
	Data      []byte
}

func main() {
	conn, err := net.Dial("tcp", ":8888")
	utils.ErrorCheck(err)

	_, err = conn.Write([]byte("input2\n"))
	utils.ErrorCheck(err)

	newFile, err := os.Create("output.png")
	utils.ErrorCheck(err)

	var protocol Protocol

	decoder := json.NewDecoder(conn)
	err = decoder.Decode(&protocol)
	utils.ErrorCheck(err)

	if protocol.WithError == true {
		fmt.Println("Response with error")
	}

	newFile.Write(protocol.Data)
}
