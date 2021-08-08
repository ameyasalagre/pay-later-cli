package modules

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ledger struct {
	User            primitive.ObjectID `bson:"user, omitempty"`
	TransactionId   string             `bson:"transaction_id" json:"transaction_id"`
	Merchant        primitive.ObjectID `bson:"merchant, omitempty"`
	Amount          float32            `json:"amount"`
	DiscountOffered float32            `json:"discount"`
	RunningBalance  float32            `bson:"running_balance" json:"running_balance"`
	LedgerType      string             `json:"type"`
	CreatedAt       primitive.Timestamp  `bson:"created_at" json:"created_at"`
	Description     string              `bson:"description" json:"description"`
}

func CreatePayout(amount, discountOffered, runningBalance float32, merchantId, user primitive.ObjectID, transactionId, ledgerType string) *Ledger {
	return &Ledger{
		User:            user,
		Merchant:        merchantId,
		TransactionId:   transactionId,
		Amount:          amount,
		DiscountOffered: discountOffered,
		RunningBalance:  runningBalance,
		LedgerType:      ledgerType,
	}
}

type MLedger struct {
	User            primitive.ObjectID `bson:"user, omitempty"`
	TransactionId   string             `bson:"transaction_id" json:"transaction_id"`
	Merchant        primitive.ObjectID `bson:"merchant, omitempty"`
	Amount          float32            `json:"amount"`
	DiscountOffered float32            `json:"discount"`
	RunningBalance  float32            `bson:"running_balance" json:"running_balance"`
	LedgerType      string             `json:"type"`
	CreatedAt       primitive.Timestamp  `bson:"created_at" json:"created_at"`
}


