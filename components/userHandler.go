package components

import (
	"context"
	"log"

	"pay-later/modules"
	"pay-later/utils"
)

func CreateUser(username, emailid string, credit, limit float32) {

	user := modules.NewUser(username, emailid,credit, limit )
	collection := utils.ConnectAndGetMongoDbCollection("pay-later", "user")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	go utils.CloseMongoDbClient()
	log.Print("User Created",username)
}


