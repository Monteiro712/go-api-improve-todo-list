package routers

import (
	"log"
	"net/http"

	"github.com/Monteiro712/go-api-improve-todo-list/controllers"
	"github.com/Monteiro712/go-api-improve-todo-list/middleware"
	"github.com/gorilla/mux"
)

func HandlerRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/tasks", controllers.GetAllTaks).Methods("Get")
	r.HandleFunc("/api/tasks/{id}", controllers.GetTask).Methods("Get")
	r.HandleFunc("/api/tasks", controllers.CreateTask).Methods("Post")
	r.HandleFunc("/api/tasks", controllers.UpdateTask).Methods("Put")
	r.HandleFunc("/api/tasks{id}", controllers.DeleteTask).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", r))
}