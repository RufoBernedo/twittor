package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/RufoBernedo/twittor/models"
)

/*Chequea si ya existe un usuario en la BD*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("usuarios")

	condicion := bson.M{"email": email} //castea a tipo BSON

	var resultado models.Usuario

	error := collection.FindOne(ctx, condicion).Decode(&resultado)

	ID := resultado.ID.Hex()
	if error != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
