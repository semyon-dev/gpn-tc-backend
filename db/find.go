package db

import (
	"context"
	"github.com/semyon-dev/gpn-tc-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func FindInUtilityModel(text string) (items []model.UtilityModel) {
	items = []model.UtilityModel{}
	filter := bson.M{"$text": bson.M{"$search": text}}
	opts := options.Find().SetSort(bson.M{"score": bson.M{"$meta": "textScore"}}).SetLimit(20).SetSkip(0)
	cursor, err := db.Collection("utility_models").Find(context.Background(), filter, opts)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &items); err != nil {
		log.Println(err)
	}
	return
}

func FindOkved(text string) (items []model.Okved) {
	items = []model.Okved{}
	filter := bson.M{"$text": bson.M{"$search": text}}
	opts := options.Find().SetSort(bson.M{"score": bson.M{"$meta": "textScore"}}).SetLimit(15).SetSkip(0)
	cursor, err := db.Collection("okveds").Find(context.Background(), filter, opts)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &items); err != nil {
		log.Println(err)
	}
	return
}
