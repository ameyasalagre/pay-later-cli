package components

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
	"log"
	"pay-later/modules"
	"pay-later/utils"
)

func generateLedger(ledger modules.Ledger) {
	collection := utils.ConnectAndGetMongoDbCollection("pay-later", "ledgers")
	
	_, err := collection.InsertOne(context.TODO(), ledger)
	if err != nil {
		log.Fatal(err)
	}
	go utils.CloseMongoDbClient()
}

func processDebitCreditFromClient(user primitive.ObjectID, runningBalance float32) string {
	
	collection := utils.ConnectAndGetMongoDbCollection("pay-later", "user")
	filter := bson.D{{"_id", user}}
	update := bson.D{{"$set",
		bson.D{
			{"credit", runningBalance},
		},
	}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
			if err == mongo.ErrNoDocuments {
			log.Println("rejected! (reason: user not found)")
			return "rejected! (reason: user not found)"
		}
	}
	go utils.CloseMongoDbClient()
	return ""
}

func processCreditToMerchant(mLedger modules.MLedger) {
	// insert new Entry in Merchant Ledger
	collection := utils.ConnectAndGetMongoDbCollection("pay-later", "merchant-ledgers")
	_, err := collection.InsertOne(context.TODO(), mLedger)
	if err != nil {
		log.Fatal(err)
	}
	mCollection := utils.ConnectAndGetMongoDbCollection("pay-later", "merchant")

	// update Merchant Running Balance
	filter := bson.D{{"_id", mLedger.Merchant}}
	update := bson.D{{"$set",
		bson.D{
			{"balance", mLedger.RunningBalance},
		},
	}}

	_, err1 := mCollection.UpdateOne(context.TODO(), filter, update)
	if err1 != nil {
		log.Fatal(err)
	}
	go utils.CloseMongoDbClient()

}

func payBackToUser(user primitive.ObjectID, runningBalance, amount float32) {
	collection := utils.ConnectAndGetMongoDbCollection("pay-later", "user")
	filter := bson.D{{"user", user}}
	update := bson.D{{"$set",
		bson.D{
			{"credit", runningBalance + amount},
		},
	}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

}
 
func getUserByUserName(userName string) (modules.User, string)  {
	var client modules.User
	collection := utils.ConnectAndGetMongoDbCollection("pay-later", "user")

	err := collection.FindOne(context.TODO(), bson.M{"username": userName}).Decode(&client)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("rejected! (reason: user not found)")
			return client,"rejected! (reason: user not found)"
		}
		log.Fatal(err)
	}
	return client,""
}

func getMerchantByName(merchantName string) (modules.Merchant, string) {
	var merchant modules.Merchant
	merchantcollection := utils.ConnectAndGetMongoDbCollection("pay-later", "merchant")
	err := merchantcollection.FindOne(context.TODO(), bson.M{"merchant_name": merchantName}).Decode(&merchant)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Merchant not found")
			return merchant,"rejected! (reason: Merchant not found)"
		}
		log.Fatal(err)
	}
	return merchant,""
}

func getAllUsers() ([]*modules.User, string)  {
	var client []*modules.User
	collection := utils.ConnectAndGetMongoDbCollection("pay-later", "user")

	cur, error := collection.Find(context.TODO(), bson.D{{}})
	if(error!=nil){
		log.Fatal(error)
		return client,error.Error()
	}

	for cur.Next(context.TODO()) {
		var resultHolder modules.User 
		error := cur.Decode(&resultHolder)
		if error != nil {
		log.Fatal(error)
		}
		client = append(client, &resultHolder)
		}
		//dont forget to close the cursor
		defer cur.Close(context.TODO())
		for _, element := range client {
			clients := *element
			log.Print(" User: " ,clients.UserName)
			log.Print(" Dues: ", float32(clients.Limit - clients.Credit) )
			}
		return client,""
}

