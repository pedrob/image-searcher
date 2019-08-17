package main

import (
	"os"

	"github.com/PedroCosta8/sistemas-distribuidos/utils"
)

func main() {
	file, err := os.Open("test")
	utils.ErrorCheck(err)
	defer file.Close()
}
