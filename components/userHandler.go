package components

import (
	"context"
	"fmt"
	"log"

	// "pay-later/modules"
	"pay-later/modules"
	"pay-later/utils"
)

func CreateUser(username, emailid string, credit, limit float32) {

	user := modules.NewUser(username, emailid,credit, limit )
	fmt.Println("Reach ")
	collection := utils.ConnectAndGetMongoDbCollection("pay-later", "user")
	data, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB", data)
	go utils.CloseMongoDbClient()
}


