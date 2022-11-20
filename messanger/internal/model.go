package internal

type Message struct {
	CreatedAt int64  `bson:"created_at" json:"created_at"`
	Text      string `bson:"text" json:"text"`
	From      string `bson:"from" json:"from"`
}
