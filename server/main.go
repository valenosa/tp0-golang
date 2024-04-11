package main

import (
	"net/http"
	"server/utils"
)

func main() {
	//Default Handler
	mux := http.NewServeMux()

	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
