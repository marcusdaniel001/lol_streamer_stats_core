package infra

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"igor.com.br/lol/lib"
)

type MongoDB struct {
	Credential *options.ClientOptions
}

func (m *MongoDB) GetAccount(account *lib.LeagueAccount, userName string) *lib.LeagueAccount {
	m.CreateConnection()
	filter := bson.M{"name": userName}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.CreateConnection().Database("leagueapi").Collection("account")
	err := collection.FindOne(ctx, filter).Decode(account)
	if err != nil {
		log.Println(err)
	}

	return account
}
func (m *MongoDB) AddNewAccount(account *lib.LeagueAccount) {
	client := m.CreateConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database("leagueapi").Collection("account")
	fmt.Println(collection)
	_, insertErr := collection.InsertOne(ctx, account)
	if insertErr != nil {
		panic(insertErr)
	} else {
		log.Println("New account added with name", account.Name)
	}
}
func (m *MongoDB) CreateConnection() *mongo.Client {
	client, err := mongo.NewClient(m.Credential)
	if err != nil {
		log.Println(err)
	}

	return client
}

func NewMongoDB(user string, password string) *MongoDB {
	return &MongoDB{
		Credential: options.Client().ApplyURI("mongodb://" + user + ":" + password + "@localhost:27017"),
	}
}
