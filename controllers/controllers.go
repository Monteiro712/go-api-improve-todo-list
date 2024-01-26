package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Monteiro712/go-api-improve-todo-list/db"
	"github.com/Monteiro712/go-api-improve-todo-list/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home page")
}

func GetAllTaks(w http.ResponseWriter, r *http.Request)  {
	var t []models.Tasks
	db.DB.Find(&t)
	json.NewEncoder(w).Encode(t)
}

func GetTask(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	var task models.Tasks
	db.DB.First(task, id)
	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request)  {
	var newTask models.Tasks
	json.NewDecoder(r.Body).Decode(&newTask)
	db.DB.Create(&newTask)
	json.NewEncoder(w).Encode(newTask)
}

func UpdateTask(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	var task models.Tasks
	db.DB.First(&task, id)
	json.NewDecoder(r.Body).Decode(&task)
	db.DB.Save(&task)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	var task models.Tasks
	db.DB.Delete(task, id)
	json.NewEncoder(w).Encode(task)
}