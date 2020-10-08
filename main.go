package main

import (
	"ApiTwitter_go/bd"
	"ApiTwitter_go/handlers"
	"log"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
	}

	handlers.Manejadores()
}
