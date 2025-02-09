package controller

type User struct {
	FName        string `json:"fName"`
	LName        string `json:"lName"`
	IdentityCard string `json:"id"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Phone        string `json:"phoneNumber"`
}
