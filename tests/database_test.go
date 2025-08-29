package tests

import (
	database "manajemen_warehouse/internal/config"
	"testing"
)

func TestDbPing(t *testing.T) {

	db := database.NewDB()
	err := db.Ping()

	if err != nil {
		// Jika err tidak nil, berarti ada yang salah.
		// GAGALKAN TES-nya dan beri pesan yang jelas.
		t.Fatalf("Expected database to ping successfully, but got an error: %v", err)
	}

}
