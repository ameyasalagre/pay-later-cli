package components

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"context"
	"fmt"
	"log"
	"strconv"
	"time"
	"pay-later/modules"
	"pay-later/utils"
)

func InitiatePayout(userName, merchant string, amount float32) string {
	// client
	var client modules.User
	client,error := getUserByUserName(userName)
	if(error!=""){
		return error
	}
	if client.Credit < amount {
		fmt.Println("Can't do transaction ,outof credit limit")
		return "rejected! (reason: credit limit)"
	}
	// Merchant

	var _merchant modules.Merchant
	_merchant,merchantErr := getMerchantByName(merchant)
	if(merchantErr!=""){
		return error
	}

	discount := _merchant.Discount
	payableForMerchant := amount - (amount / 100 * discount)
	transactionId := time.Now().UnixNano()
	// ledger
	runningBalance := client.Credit - amount


	timeStamp := primitive.Timestamp{T:uint32(time.Now().Unix())}
	var ledger modules.Ledger
	ledger.Amount = amount
	ledger.DiscountOffered = _merchant.Discount
	ledger.LedgerType = "DEBIT"
	ledger.Merchant = _merchant.Id
	ledger.TransactionId = "PL" + strconv.Itoa(int(transactionId))
	ledger.User = client.Id
	ledger.RunningBalance = runningBalance
	ledger.CreatedAt = timeStamp
	processDebitCreditFromClient(client.Id,runningBalance)

	var mLedger modules.MLedger
	mLedger.Amount = payableForMerchant
	mLedger.CreatedAt = timeStamp
	mLedger.DiscountOffered = _merchant.Discount
	mLedger.LedgerType = "CREDIT"
	mLedger.Merchant = _merchant.Id
	mLedger.RunningBalance = _merchant.Balance + payableForMerchant
	mLedger.TransactionId = "PL" + strconv.Itoa(int(transactionId))
	mLedger.User = client.Id
	processCreditToMerchant(mLedger)

	
	generateLedger(ledger)
	fmt.Println(transactionId)

	defer utils.CloseMongoDbClient()
	return "trassssss"

}

func createPayout(amount, discountOffered, runningBalance float32, merchant, user primitive.ObjectID, transactionId, ledgerType string) {
	ledgerEntry := modules.CreatePayout(amount, discountOffered, runningBalance, merchant, user, transactionId, ledgerType)
	fmt.Println("Reach ", ledgerEntry)
	collection := utils.ConnectAndGetMongoDbCollection("pay-later", "ledgers")
	data, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB", data)
	go utils.CloseMongoDbClient()
}

func CreatePayBack(userName string, amount float32) string {
	var client modules.User
	var ledger modules.Ledger
	transactionId := time.Now().UnixNano()
	timeStamp := primitive.Timestamp{T:uint32(time.Now().Unix())}

	client,error :=getUserByUserName(userName)
	if(error!="") {
		return error
	}
	// client.Limit // 
	due := client.Limit - client.Credit
	restoredCredit := client.Credit + amount

	if(restoredCredit  > client.Limit){
		fmt.Println("Cannot pay beyond limit,maximum payable amount is", float64(due))
		return fmt.Sprint("Cannot pay beyond limit,maximum payable amount is", float64(due))
	}
	
	processDebitCreditFromClient(client.Id,restoredCredit)

	ledger.Amount = amount
	ledger.LedgerType = "CREDIT"
	ledger.TransactionId = "PB" + strconv.Itoa(int(transactionId))
	ledger.User = client.Id
	ledger.RunningBalance = restoredCredit
	ledger.CreatedAt = timeStamp
	ledger.Description = "Payback by user"
	generateLedger(ledger)
	
	return fmt.Sprint("Payback Successfull for amount", float64(due))
}
