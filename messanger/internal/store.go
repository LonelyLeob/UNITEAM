package internal

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx = context.TODO()
)

type Mongoose struct {
	config *ConfigMongo
	client *mongo.Client
}

func InitMongoose(cfg *ConfigMongo) *Mongoose {
	return &Mongoose{
		config: cfg,
	}
}

func (m *Mongoose) Connect() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:root@mongoservice:27017"))
	if err != nil {
		return err
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return err
	}

	m.client = client

	return nil
}

func (m *Mongoose) CreateMessage(msg *Message) error {
	coll := m.client.Database("messanger").Collection("messages")

	_, err := coll.InsertOne(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongoose) GetMessagesWithOffset(limit int64) ([]*Message, error) {
	var msgs []*Message
	coll := m.client.Database("messanger").Collection("messages")
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"created_at", -1}}).SetLimit(limit)
	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		msg := &Message{}
		if err := cursor.Decode(msg); err != nil {
			fmt.Println(err)
			continue
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}
