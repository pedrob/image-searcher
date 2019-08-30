package tcp

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"

	"github.com/PedroCosta8/sistemas-distribuidos/utils"
)

func SearchImage(imageName string) (*os.File, error) {
	conn, err := net.Dial("tcp", ":8888")
	defer conn.Close()
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

	fileSize, err := io.Copy(image, conn)
	if err != nil {
		return nil, err
	}

	if fileSize == 0 {
		err = utils.ImageNotFound()
		return nil, err
	}
	return image, nil
}
