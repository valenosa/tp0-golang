package main

import (
	"client/globals"
	"client/utils"
)

func main() {
	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	// validar que la config este cargada correctamente

	// loggeamos el valor de la config

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config

	// leer de la consola el mensaje
	// utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	// utils.GenerarYEnviarPaquete()
}
