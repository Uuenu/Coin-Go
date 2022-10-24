package mongodb

import (
	"context"
	"fmt"
	"strconv"
	lib "telegram-coin-go/lib/e"
	"telegram-coin-go/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbURI = "mongodb://localhost:27017"
)

type Storage struct {
	Client *mongo.Client
	DB     *mongo.Database
	Ctx    context.Context
}

func New() *Storage {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil
	}
	database := client.Database("Coin-Go")
	if err != nil {
		return nil
	}
	s := Storage{
		Client: client,
		DB:     database,
		Ctx:    context.TODO(),
	}

	return &s
}

func (s Storage) AddRecord(page *storage.Page) (err error) {
	defer func() { err = lib.WrapIfErr("can't save page", err) }()
	userCollection := s.DB.Collection(strconv.Itoa(page.ChatID))
	doc, err := DataToBson(page.Data)
	result, err := userCollection.InsertOne(context.TODO(), doc)

	fmt.Println(&result)

	return err
}

func (s Storage) RecordsList(chatID int, limit int) ([]storage.Page, error) {
	userCollection := s.DB.Collection(strconv.Itoa(chatID))
	records := make([]*mongo.Cursor, 0)

	for i := 0; i < limit; i++ {
		result, err := userCollection.Find(context.TODO(), bson.M{"record_id": i}) // last id - i
		if err != nil {
			return nil, err
		}
		records = append(records, result)
	}

	// []

	return nil, nil
}

func DataToBson(data map[string]string) (bson.D, error) {
	return nil, nil
}
