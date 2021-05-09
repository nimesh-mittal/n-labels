package gateway

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// TODO: check if below import is really required
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

// MongoClient represents mongo client apis
type MongoClient interface {
	GetDocByID(db string, col string, result interface{}, field string, value interface{}) error
	DeleteDocByID(db string, col string, filter map[string]interface{}) (bool, error)
	ListDocs(db string, col string, results interface{}, filter map[string]interface{}, limit int64, offset int64) error
	InsertDoc(db string, col string, doc interface{}) error
	UpdateDocByID(db string, col string, field string, value interface{}, updateKey string, updateValue interface{}) (bool, error)
	Close()
}

type mongoClient struct {
	Client *mongo.Client
}

// New creates new object of MongoClient
func New(url string) MongoClient {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		// TODO: handle this error
	}

	return &mongoClient{Client: client}
}

func (mc *mongoClient) Close() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mc.Client.Disconnect(ctx)
}

func (mc *mongoClient) ListDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	databases, err := mc.Client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(databases)
}

func (mc *mongoClient) GetDocByID(db string, col string, result interface{}, field string, value interface{}) error {
	database := mc.Client.Database(db)
	collection := database.Collection(col)

	filter := bson.D{{Key: field, Value: value}}

	if field == "" {
		filter = bson.D{}
	}

	err := collection.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (mc *mongoClient) DeleteDocByID(db string, col string, filter map[string]interface{}) (bool, error) {
	database := mc.Client.Database(db)
	collection := database.Collection(col)

	f := bson.D{}
	for k, v := range filter {
		f = append(f, bson.E{Key: k, Value: v})
	}

	del, err := collection.DeleteOne(context.TODO(), f)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return del.DeletedCount == 1, nil
}

func (mc *mongoClient) ListDocs(db string, col string, results interface{}, filter map[string]interface{}, limit int64, offset int64) error {
	database := mc.Client.Database(db)
	collection := database.Collection(col)

	f := bson.D{}
	for k, v := range filter {
		f = append(f, bson.E{Key: k, Value: v})
	}

	log.Println(f)

	op := options.Find()
	op.SetSkip(offset)
	op.SetLimit(limit)

	cursor, err := collection.Find(context.TODO(), f, op)
	if err != nil {
		log.Println(err)
		return err
	}

	cursor.All(context.TODO(), results)

	return nil
}

func (mc *mongoClient) InsertDoc(db string, col string, doc interface{}) error {
	database := mc.Client.Database(db)
	collection := database.Collection(col)

	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (mc *mongoClient) UpdateDocByID(db string, col string, field string, value interface{}, updateKey string, updateValue interface{}) (bool, error) {
	database := mc.Client.Database(db)
	collection := database.Collection(col)

	filter := bson.M{field: value}
	updatedDoc := bson.D{{"$set", bson.D{{updateKey, updateValue}}}}

	if field == "" {
		filter = bson.M{}
	}

	res, err := collection.UpdateOne(context.TODO(), filter, updatedDoc)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return res.ModifiedCount == 1, nil
}
