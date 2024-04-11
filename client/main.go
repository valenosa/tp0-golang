package main

import (
	"client/globals"
	"client/utils"
)

func main() {

	//Establece un output dual hacia consola y "tp0.log"
	utils.ConfigurarLogger()

	//Creo un paquete
	paquete := utils.Paquete{}

	//Extrae la informacion de config.json y la almacena en una variable llamada ClientConfig contenida dentro de globals
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// Se envia el mensaje
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	// Recibe logs por consola y los va almacenando en un paquete-buffer. Una vez se termina el programa (al apretar enter sin escribir nada), sube los logs de este paquete-buffer al paquete original
	utils.ManejarPaquete(&paquete)

	// Envia el paquete en el cuerpo de una Request
	utils.EnviarPaquete(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, paquete)
}
