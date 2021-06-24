package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

// "Person type" (tipo um objeto)
type Message struct {
    Mensagem  string `json:"mensagem"`
}

var resposta []Message

func GetResposta(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(resposta)
}

func main() {
    router := mux.NewRouter()
	resposta = append(resposta, Message{Mensagem: "Hello, World!"})
    router.HandleFunc("/", GetResposta).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}