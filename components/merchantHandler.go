package components

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
	"log"
	"pay-later/modules"
	"pay-later/utils"
)

func CreateMerchant(merchantName string, discount float32) {
	merchant := modules.NewMerchant(merchantName, discount)
	collection := utils.ConnectAndGetMongoDbCollection("pay-later", "merchant")
	_, err := collection.InsertOne(context.TODO(), merchant)
	if err != nil {
		log.Fatal(err)
	}
	go utils.CloseMongoDbClient()
	log.Print("Merchant Created Successfully")
}

func UpDateMerchantDiscountByName(merchantName string, updatedDiscount float32) {
	collection := utils.ConnectAndGetMongoDbCollection("pay-later", "merchant")

	filter := bson.D{{"merchant_name", merchantName}}
	update := bson.D{{"$set",
		bson.D{
			{"discount", updatedDiscount},
		},
	}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("rejected! (reason: user not found)")
		}
	}
	go utils.CloseMongoDbClient()
	log.Println("Merchant Discount updated Successfully")
}
