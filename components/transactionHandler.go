package components

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

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
		log.Print("Can't do transaction ,out of credit limit rejected! (reason: credit limit)")
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
	
	processDebitCreditFromClient(client.Id,runningBalance)
	//map merchant ledger
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

	//map ledger struct
	var ledger modules.Ledger
	ledger.Amount = amount
	ledger.DiscountOffered = _merchant.Discount
	ledger.LedgerType = "DEBIT"
	ledger.Merchant = _merchant.Id
	ledger.TransactionId = "PL" + strconv.Itoa(int(transactionId))
	ledger.User = client.Id
	ledger.RunningBalance = runningBalance
	ledger.CreatedAt = timeStamp
	generateLedger(ledger)
	log.Print("Transaction Success Transaction-ID PL", strconv.Itoa(int(transactionId)))
	defer utils.CloseMongoDbClient()
	return "Transaction Success!"

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
		log.Println("Cannot pay beyond limit,maximum payable amount is", float64(due))
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
	log.Print("Payback Successfull TransactionId: ", "PB" + strconv.Itoa(int(transactionId)))
	log.Println("(Dues)",client.Limit-restoredCredit)

	return fmt.Sprint("Payback Successfull for amount", float64(due))
}
