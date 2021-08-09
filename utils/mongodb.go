package utils

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/****************************************.
*	MongoDB getConnection
* 	MongoDb getCollection
* 	MongoDb  closeConnection
*  	TODO: Logging,Refactoring
******************************************/

var client *mongo.Client

/********
*	connect To MongoDb and get Collection by passing documentName,tableName as a String
* 	and Return the collection pointer
*********/
func ConnectAndGetMongoDbCollection(documentName string, tableName string) *mongo.Collection {
	// get a MongoDBClient
	client := getClient()
	// Check the connection
	collection := client.Database(documentName).Collection(tableName)
	return collection
}

/********
*	Close  MongoDb connection
*	TODO:
*	~Stop MongoClient when its ideal for more than 2 mins
*********/
func CloseMongoDbClient() {
	client := getClient()
	client.Disconnect(context.TODO())
}


/********
*	get MongoDb client
*********/
func getClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:8Qo3iWUE68khYBo6@golang-playground.hsksi.mongodb.net/playground?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error While Connection", err)
	}
	return client 
}




