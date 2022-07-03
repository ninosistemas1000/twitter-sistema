package main

import (
	"log"

	"github.com/ninosistemas1000/twitter-sistema/bd"
	"github.com/ninosistemas1000/twitter-sistema/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
