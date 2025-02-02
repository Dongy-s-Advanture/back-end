package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Seller struct {
	SellerID    primitive.ObjectID `json:"sellerID"`
	Username    string             `json:"username"`
	Name        string             `json:"name"`
	Surname     string             `json:"surname"`
	Payment     string             `json:"payment"`
	Address     string             `json:"address"`
	PhoneNumber string             `json:"phoneNumber"`
	Score       float32            `json:"score"`
}

type SellerRegisterRequest struct {
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	Name        string  `json:"name"`
	Surname     string  `json:"surname"`
	Payment     string  `json:"payment"`
	Address     string  `json:"address"`
	PhoneNumber string  `json:"phoneNumber"`
	Score       float32 `json:"score"`
}
