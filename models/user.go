package models

type Address struct{
	State string	`json:"state" bson:"state"`
	City string		`json:"city" bson:"city"`
	Pincode int		`json:"pincode" bson:"pincode"`
}

type Date struct{
	Day 	int `json:"day" bson:"day"`
	Month 	int	`json:"month" bson:"month"`
	Year 	int	`json:"year" bson:"year"`
}

type User struct{
	Id 			string		`json:"id" bson:"id"`
	Name 		string	`json:"name" bson:"name"`
	Dob 		Date	`json:"dob" bson:"dob"`
	Address 	Address	`json:"address" bson:"address"`
	Description string	`json:"description" bson:"description"`
	CreatedAt 	Date 	`json:"createdAt" bson:"createdAt"`
}
