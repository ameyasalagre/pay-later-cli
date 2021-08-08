package modules

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	UserName string             `json:"username" `
	Credit   float32            `json:"credit" `
	Email    string             `json:"email"`
	Limit    float32             `json:"limit"`

}

func NewUser(userName ,email string, limit, credit float32) *User {
	return &User{
		UserName: userName,
		Credit:   credit,
		Email:    email,
		Limit: limit,
	}
}

func (u *User) GetId() string {
	return u.UserName
}

func (u User) Data(name string) string {
	u.UserName = "Father : " + name
	return u.UserName
}
