package main

import (
	"fmt"

	"github.com/Monteiro712/go-api-improve-todo-list/db"
)

func main() {
	db.ConectarBancoDeDados()
	fmt.Println("iniciando")
}