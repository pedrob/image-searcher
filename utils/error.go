package utils

import (
	"errors"
	"fmt"
	"os"
)

func ErrorCheck(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
	}
}

func ImageNotFound() error {
	return errors.New("Image not found in the server")
}
