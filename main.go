package main

import (
	"log"

	"github.com/Greyes93/twittor/bd"
	"github.com/Greyes93/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
	}

	handlers.Manejadores()
}
