package main

import "github.com/ShivanshuPrajapati212/ascii-football-server/pkg/database"

func main() {
	database.ConnectToMongoDB("mongodb://localhost:27017")
}
