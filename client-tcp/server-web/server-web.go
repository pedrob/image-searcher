package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./client-tcp/assets")))
	//http.HandleFunc("/send", uploadFileAndSend)
	http.ListenAndServe(":3000", nil)
}

// func uploadFileAndSend(w http.ResponseWriter, r *http.Request) {
// 	//valor maximo do upload
// 	r.ParseMultipartForm(10 << 20)

// 	//recebe o arquivo vindo do formulario
// 	file, handler, err := r.FormFile("file")
// 	utils.ErrorCheck(err)
// 	defer file.Close()

// 	fmt.Println(handler.Filename)

// 	// converte o arquivo em um array de bytes
// 	fileBytes, err := ioutil.ReadAll(file)
// 	utils.ErrorCheck(err)

// 	//envia o array de bytes para o servidor TCP
// 	err = cliTcp.TcpSendFile(fileBytes)
// 	utils.ErrorCheck(err)
// }
