package main

import (
	"fmt"

	"api-rest-1/database"
	"api-rest-1/models"
	"api-rest-1/routes"

)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectComBancoDoDados()
	fmt.Println("Iniciando o Servido")
	routes.HandleRequest()
}
