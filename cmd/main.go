package main

import (
	database "manajemen_warehouse/internal/config"
)

func main() {
	db := database.NewDB()
	db.Ping()
}
