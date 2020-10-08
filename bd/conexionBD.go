package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la Base de Datos */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://administrador:00JtUIuW6BCAApIp@apicluster.gcfws.mongodb.net/twittor?retryWrites=true&w=majority")

/*ConectarBD es la función que me permite conectar la Base de Datos */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Conexión Exitosa con la BD")
	return client
}

/*ChequeoConnection es el ping a la Base de Datos*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
