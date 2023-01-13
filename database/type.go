package database

import "gorm.io/gorm"

// Users struct is used to store the data from the user to database
type Users struct {
	gorm.Model
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// Login struct is used to login in the system
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Updation struct is used for the updation
type Updations struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Details struct is used for the listing of the users
type Details struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Response struct is used for the response
type Response struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
