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

func (s Storage) AddRecord(record *storage.Record) (err error) {
	defer func() { err = lib.WrapIfErr("can't save page", err) }()
	userCollection := s.DB.Collection(strconv.Itoa(record.ChatID))
	result, err := userCollection.InsertOne(context.TODO(), record)
	fmt.Println(&result)
	return err
}

func (s Storage) UpdateLastRecord(chatID int, sum float64) (err error) {
	userCollection := s.DB.Collection(strconv.Itoa(chatID))
	lastRec, _ := s.LastRecord(chatID)
	lastSum := lastRec.Day.Total

	update := bson.D{
		{Key: "day", Value: bson.D{
			{Key: "sum", Value: lastSum + sum},
		}},
	}

	opts := options.FindOneAndUpdate().SetSort(bson.M{"$natural": -1})
	userCollection.FindOneAndUpdate(context.TODO(), bson.M{}, bson.M{"$set": update}, opts)
	return nil
}

func (s Storage) DaysList(chatID int, limit int) ([]storage.Record, error) {

	return nil, nil
}

func (s Storage) LastRecord(chatID int) (result storage.Record, err error) {

	userCollection := s.DB.Collection(strconv.Itoa(chatID))

	opts := options.FindOne().SetSort(bson.M{"$natural": -1})

	if err = userCollection.FindOne(context.TODO(), bson.M{}, opts).Decode(&result); err != nil {
		return storage.Record{}, err
	}

	return result, err

}

func (s Storage) Today(chatID int) (storage.Record, error) {
	return s.LastRecord(chatID)
}

func (s Storage) CheckTime(ChatID int, TimeNow string) bool {
	lastRec, _ := s.LastRecord(ChatID)
	return lastRec.Time == TimeNow
}
