package main

import (
	"github.com/onasunnymorning/go-gin-boilerplate/model"
	"github.com/onasunnymorning/go-gin-boilerplate/route"
)

func main() {

	db, _ := model.DBConnection()
	route.SetupRoutes(db)
}