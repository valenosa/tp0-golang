package main

import (
	"client/globals"
	"client/utils"

	//No quiero tener que importar log, quiero que con importar utils alcance
	"log"
)

func main() {
	utils.ConfigurarLogger()

	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}

	// loggeamos el valor de la config
	log.Println(globals.ClientConfig.Mensaje)

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config

	// leer de la consola el mensaje
	// utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	// utils.GenerarYEnviarPaquete()
}
