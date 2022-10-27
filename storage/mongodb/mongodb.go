package mongodb

import (
	"context"
	"fmt"
	"strconv"
	lib "telegram-coin-go/lib/e"
	"telegram-coin-go/storage"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// get last record
	data, _, _ := s.LastRecord(chatID)

	lastSum, _ := strconv.ParseFloat(data["Sum"], 64)

	nwRecSum, _ := strconv.ParseFloat(record["Sum"], 64)

	updateSum := lastSum + nwRecSum
	// calc new sum

	// update
	update := bson.D{
		{Key: "Data", Value: bson.D{
			//{Key: "Debit", Value: record["debit"]},
			//{Key: "Credit", Value: record["credit"]},
			{Key: "Sum", Value: strconv.FormatFloat(updateSum, 'f', 5, 64)},
			//{Key: "Text", Value: record.Data["text"]},
		}},
	}

	opts := options.FindOneAndUpdate().SetSort(bson.M{"$natural": -1})
	userCollection.FindOneAndUpdate(context.TODO(), bson.M{}, update, opts)

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

func (s Storage) LastRecord(chatID int) (map[string]string, time.Time, error) {

	userCollection := s.DB.Collection(strconv.Itoa(chatID))

	opts := options.FindOne().SetSort(bson.M{"$natural": -1})
	var lastrecord bson.M
	if err := userCollection.FindOne(context.TODO(), bson.M{}, opts).Decode(&lastrecord); err != nil {
		return nil, time.Time{}, err
	}

	data := lastrecord["Data"].(primitive.M)
	result := make(map[string]string)
	for key, value := range data {
		result[key] = fmt.Sprintf("%s", value)
	}

	recTime := lastrecord["Time"].(time.Time)

	fmt.Println("data + time ", result, " ", recTime, " ")
	return result, recTime, nil

}

func RecordToBson(record *storage.Record) (bson.D, error) {
	result := bson.D{
		{Key: "ChatID", Value: record.ChatID},
		{Key: "Username", Value: record.Username},
		{Key: "Time", Value: record.Time},
		{Key: "Data", Value: bson.D{
			//{Key: "Debit", Value: record.Data["debit"]},
			//{Key: "Credit", Value: record.Data["credit"]},
			{Key: "Sum", Value: record.Data["Sum"]},
			//{Key: "Text", Value: record.Data["text"]},
		}},
	}
	return result, nil
}
