package tcp

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
)

type Response struct {
	WithError bool
	Data      []byte
}

func SearchImage(imageName string) (*os.File, error) {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		return nil, err
	}

	_, err = conn.Write([]byte(fmt.Sprintf("%s\n", imageName)))
	if err != nil {
		return nil, err
	}

	image, err := ioutil.TempFile(os.TempDir(), "image.png")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(image, conn)
	if err != nil {
		return nil, err
	}
	return image, nil
}
