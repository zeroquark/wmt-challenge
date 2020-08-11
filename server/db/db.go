package db

import (
	"context"
	"log"
	"walmart/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	FindByID(id int) (*model.Product, error)
	FindByDescription(description string) ([]*model.Product, error)
	FindByBrand(brand string) ([]*model.Product, error)
}

type MongoDB struct {
	collection *mongo.Collection
}

func NewMongo(client *mongo.Client) DB {
	database := "promotions"
	collection := "products"
	tech := client.Database(database).Collection(collection)
	return MongoDB{collection: tech}
}

// Find product corresponding to ID
func (m MongoDB) FindByID(id int) (*model.Product, error) {
	var result model.Product
	filter := bson.D{{Key: "id", Value: id}}

	err := m.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil && result.Brand == "" {
		return nil, err
	}

	return &result, nil
}

// Find products by brand
func (m MongoDB) FindByBrand(brand string) ([]*model.Product, error) {
	// Collection of found documents (Product)
	var results []*model.Product
	// Set filter and find options
	filter := bson.D{{Key: "brand", Value: primitive.Regex{Pattern: brand}}}
	// findOptions := options.Find().SetLimit(3)
	// Set the cursor
	// cur, err := db.GetDefaultCollection().Find(context.TODO(), filter, findOptions)
	cur, err := m.collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var p model.Product
		err := cur.Decode(&p)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &p)
	}

	if err := cur.Err(); err != nil {
		log.Fatal()
		return results, err
	}
	cur.Close(context.TODO())

	return results, nil
}

// Find products by description
func (m MongoDB) FindByDescription(description string) ([]*model.Product, error) {
	// Collection of found documents (Product)
	var results []*model.Product
	// Set filter and find options
	filter := bson.D{{Key: "description", Value: primitive.Regex{Pattern: description}}}
	// findOptions := options.Find().SetLimit(3)
	// Set the cursor
	// cur, err := db.GetDefaultCollection().Find(context.TODO(), filter, findOptions)
	cur, err := m.collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var p model.Product
		err := cur.Decode(&p)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &p)
	}

	if err := cur.Err(); err != nil {
		log.Fatal()
		return results, err
	}
	cur.Close(context.TODO())

	return results, nil
}
