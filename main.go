package main

import (
	"api_rest/db"
	"api_rest/routes"
	"fmt"
)

func main() {
	db.DbConnect()
	fmt.Println("Server running in :3030")
	routes.HandleRequest()
}
