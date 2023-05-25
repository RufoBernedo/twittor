package main

import (
	"log"

	"github.com/RufoBernedo/twittor/bd"
	"github.com/RufoBernedo/twittor/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
