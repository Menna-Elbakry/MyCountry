package main

import (
	database "task/database/implementation"
	"task/routes"
)

func main() {
	database.CreateTables()
	routes.Server()

}
