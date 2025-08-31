package main

import database "manajemen_warehouse/pkg/config"

func main() {
	db := database.NewDB()
	db.Ping()
}
