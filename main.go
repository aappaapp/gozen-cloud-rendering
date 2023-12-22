package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

func HandleRender(a http.ResponseWriter, b *http.Request) {
	var err error

	if b.Method != "PUT" {
		a.WriteHeader(http.StatusBadRequest)
		return
	}

	tempDir, err := os.MkdirTemp("", "gcr")

	command := exec.Command("gozen", "--render", "<PATH>")

	err = command.Run()
	if err != nil {
		log.Fatalln(err)
	}

	err = os.RemoveAll(tempDir)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/render", HandleRender)
	http.ListenAndServe(":8080", nil)
}
