package utils

import (
	"bufio"
	"bytes"
	"client/globals"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func IniciarConfiguracion(filePath string) *globals.Config {
	var config *globals.Config
	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer configFile.Close()

	//Por que no asi: jsonParser := json.NewDecoder(configFile).Decode(&config)
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

// Lee la consola y genera un paquete
func ManejarPaquete(paqueteOriginal *Paquete) {
	// Dereferencio al puntero
	paquete := *paqueteOriginal
	// Leer de la consola
	reader := bufio.NewReader(os.Stdin)

	log.Println("Ingrese los mensajes")

	//Bucle para recibir varios mensajes
	for {
		text, err := reader.ReadString('\n')

		// Manejar el error
		if err != nil {
			log.Printf("Input data error:%s", err.Error())
		}

		//Una vez presionado el caracter finalizador
		if text == "\n" {
			// Asigno el valor final al puntero (referencio)
			*paqueteOriginal = paquete

			// Genero log sobre el paquete a enviar
			log.Printf("paqute a enviar: %+v", paquete)

			break
		}
		// Genero log con el mensaje
		log.Print(text)

		// Agrego el mensaje al paquete
		paquete.Valores = append(paquete.Valores, text)
	}

}

func EnviarMensaje(ip string, puerto int, mensajeTxt string) {
	mensaje := Mensaje{Mensaje: mensajeTxt}
	body, err := json.Marshal(mensaje)
	if err != nil {
		log.Printf("error codificando mensaje: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/mensaje", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensaje a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func EnviarPaquete(ip string, puerto int, paquete Paquete) {
	body, err := json.Marshal(paquete)
	if err != nil {
		log.Printf("error codificando mensajes: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/paquetes", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensajes a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func ConfigurarLogger() {
	logFile, err := os.OpenFile("tp0.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
