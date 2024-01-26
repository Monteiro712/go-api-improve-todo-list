package models

type Tasks struct{
	ID int `json:"id"`
	Title string `json:"title"`
	About string `json:"About"`
	Completed bool `json:"completed"`
	UserID int `json:"userid"`
}

var Task []Tasks