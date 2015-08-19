package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"

	"net/http"
)

// START OMIT
func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("respond:", err)
	}
}

// END OMIT

type msg struct {
	Message string `json:"msg"`
}

func main() {
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/goodbye", handleBye)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalln(err)
	}
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	respond(w, r, http.StatusOK, &msg{Message: "Hello Golang UK Conf"})
}

func handleBye(w http.ResponseWriter, r *http.Request) {
	respond(w, r, http.StatusOK, &msg{Message: "See you next year"})

}
