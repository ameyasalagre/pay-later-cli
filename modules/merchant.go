package modules

import "go.mongodb.org/mongo-driver/bson/primitive"

type Merchant struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	MerchantName string   `bson:"merchant_name" json:"name"`
	Discount     float32  `json:"discount"`
	Balance      float32  
}

func NewMerchant(name string, discount float32) *Merchant {
	return &Merchant{
		MerchantName: name,
		Discount:     discount,
		Balance: float32(0.0),
	}
}

func (m *Merchant) GetId() string {
	return m.MerchantName
}
