package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/RufoBernedo/twittor/models"
)

func InsertoTweet(t models.GraboTweet)(string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("tweet")

	registro := bson.M{
		"userid": t.UserID,
		"mensaje": t.Mensaje,
		"fecha": t.Fecha,
	}

	result, err := collection.InsertOne(ctx, registro)
	if err != nil {
		return "",false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID) // se obtiene el ObjectId del registro insertado
	return objID.String(), true, nil
}