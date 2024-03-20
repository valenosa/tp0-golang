package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func RecibirPaquetes(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var paquete Paquete
	err := decoder.Decode(&paquete)
	if err != nil {
		log.Printf("error al decodificar mensaje: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error al decodificar mensaje"))
		return
	}

	log.Println("me llego un paquete de un cliente")
	log.Printf("%+v\n", paquete)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func RecibirMensaje(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var mensaje Mensaje
	err := decoder.Decode(&mensaje)
	if err != nil {
		log.Printf("Error al decodificar mensaje: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al decodificar mensaje"))
		return
	}

	log.Println("Me llego un mensaje de un cliente")
	log.Printf("%+v\n", mensaje)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
