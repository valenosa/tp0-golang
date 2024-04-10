package main

import (
	"client/utils"

	//No quiero tener que importar log, quiero que con importar utils alcance
	"bufio"
	"log"
	"os"
)

func main() {

	//Establece un output dual hacia consola y "tp0.log"
	utils.ConfigurarLogger()

	reader := bufio.NewReader(os.Stdin)

	for {

		text, err := reader.ReadString('\n')

		if err != nil {
			// Manejar el error
			log.Printf("Input data error:%s", err.Error())
		}

		if text == "\n" {
			break
		}

		log.Print(text)
	}

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config

	// leer de la consola el mensaje
	// utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	// utils.GenerarYEnviarPaquete()
}
