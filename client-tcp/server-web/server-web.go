package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	cliTcp "github.com/PedroCosta8/sistemas-distribuidos/client-tcp"
	"github.com/PedroCosta8/sistemas-distribuidos/utils"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./client-tcp/assets")))
	http.HandleFunc("/send", uploadFileAndSend)
	http.ListenAndServe(":3000", nil)
}

func uploadFileAndSend(w http.ResponseWriter, r *http.Request) {
	//valor maximo do upload
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	utils.ErrorCheck(err)
	defer file.Close()

	fmt.Println(handler.Filename)

	fileBytes, err := ioutil.ReadAll(file)
	utils.ErrorCheck(err)

	err = cliTcp.TcpSendFile(fileBytes)
	utils.ErrorCheck(err)
}
