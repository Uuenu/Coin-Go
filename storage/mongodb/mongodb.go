package mongodb

import (
	"context"
	"fmt"
	"strconv"
	lib "telegram-coin-go/lib/e"
	"telegram-coin-go/storage"
	"time"

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

func (s Storage) AddRecord(page *storage.Record) (err error) {
	defer func() { err = lib.WrapIfErr("can't save page", err) }()
	userCollection := s.DB.Collection(strconv.Itoa(page.ChatID))
	doc, err := RecordToBson(page)
	result, err := userCollection.InsertOne(context.TODO(), doc)

	fmt.Println(&result)

	return err
}

func (s Storage) UpdateLastRecord(chatID int, record map[string]string) (err error) {
	userCollection := s.DB.Collection(strconv.Itoa(chatID))

	update := bson.D{
		{Key: "Data", Value: bson.D{
			{Key: "Debit", Value: record["debit"]},
			{Key: "Credit", Value: record["credit"]},
			{Key: "Sum", Value: record["sum"]},
			//{Key: "Text", Value: record.Data["text"]},
		}},
	}

	userCollection.FindOneAndUpdate(context.TODO(), bson.M{"$natural": -1}, update)

	return nil
}

func (s Storage) RecordsList(chatID int, limit int) ([]storage.Record, error) {
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

func (s Storage) LastRecord(chatID int) (string, string, time.Time, error) {
	userCollection := s.DB.Collection(strconv.Itoa(chatID))
	var record bson.M
	err := userCollection.FindOne(context.TODO(), bson.M{"$natural": -1}).Decode(&record) // !!!
	if err != nil {
		return "", "", time.Now(), err
	}
	data := record["Data"].(map[string]string) // interface to map
	debit := data["Debit"]
	credit := data["Credit"]
	recTime := record["Time"].(time.Time)

	return debit, credit, recTime, nil
}

func RecordToBson(record *storage.Record) (bson.D, error) {
	result := bson.D{
		{Key: "ChatID", Value: record.ChatID},
		{Key: "Username", Value: record.Username},
		{Key: "Time", Value: record.Time},
		{Key: "Data", Value: bson.D{
			{Key: "Debit", Value: record.Data["debit"]},
			{Key: "Credit", Value: record.Data["credit"]},
			{Key: "Sum", Value: record.Data["sum"]},
			//{Key: "Text", Value: record.Data["text"]},
		}},
	}
	return result, nil
}
