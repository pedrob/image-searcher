package tcp

import (
	"net"

	"github.com/PedroCosta8/sistemas-distribuidos/utils"
)

func TcpSendFile(file []byte) error {
	connection, err := net.Dial("tcp", ":8888")
	utils.ErrorCheck(err)
	defer connection.Close()

	_, err = connection.Write(file)
	if err != nil {
		return err
	}
	return nil
}
