package main

import (
	"github.com/wellrcosta/Api.Crud.Go.Gin/database"
	"github.com/wellrcosta/Api.Crud.Go.Gin/routes"
)

func main() {
	database.ConnectToDatabase()
	err := routes.HandleRequests()
	if err != nil {
		return
	}
}
