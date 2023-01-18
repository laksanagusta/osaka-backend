package main

import (
	"tokokecilkita-go/database"
	"tokokecilkita-go/routes"
)

func main() {
	db := database.SetupDatabase()
	r := routes.SetupRoutes(db)
	r.Run()
}