package mongodb_adapter

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

type MongoClient struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

var once sync.Once
var conn *mongo.Client

func ConnectMongo(mongoUrl string, dbName, collectionName string) *MongoClient {
	var err error

	once.Do(func() {
		conn, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))

		if err != nil {
			log.Fatal(err)
		}
	})

	collection := conn.Database(dbName).Collection(collectionName)

	return &MongoClient{Client: conn, Collection: collection}
}

func (c *MongoClient) GetData(filter bson.D, returnData interface{}) error {
	err := c.Collection.FindOne(context.TODO(), filter).Decode(returnData)

	if err != nil {
		err = errors.New("Document Not Found")
	}

	return err
}

func (c *MongoClient) GetAll(filter bson.D, returnData interface{}) error {

	findOptions := options.Find()
	findOptions.SetLimit(100)
	cursor, err := c.Collection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		err = errors.New("Document Not Found:" + err.Error())
	} else {
		err = cursor.All(context.TODO(), returnData)

		if err != nil {
			err = errors.New("Document Cursor")
		}
	}
	return err
}

func (c *MongoClient) AddData(insert interface{}) error {
	_, err := c.Collection.InsertOne(context.TODO(), insert)

	if err != nil {
		return err
	}

	return nil
}

func (c *MongoClient) UpdateData(filter bson.D, update bson.D) error {
	_, err := c.Collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (c *MongoClient) DeleteData(filter bson.D) error {
	options := options.Delete()

	_, err := c.Collection.DeleteOne(context.TODO(), filter, options)

	if err != nil {
		return err
	}

	return nil
}
