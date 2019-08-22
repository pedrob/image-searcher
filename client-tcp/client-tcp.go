package main

import (
	"io"
	"net"
	"os"

	"github.com/PedroCosta8/sistemas-distribuidos/utils"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	utils.ErrorCheck(err)
	newFile, err := os.Create("output.png")
	utils.ErrorCheck(err)
	_, err = io.Copy(newFile, conn)
	utils.ErrorCheck(err)
}
