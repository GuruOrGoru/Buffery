package db

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestInit(t *testing.T) {
	// Test with invalid DSN
	_, err := Init("invalid_dsn")
	if err == nil {
		t.Fatal("Init should have failed with invalid DSN")
	}

	// Note: Successful initialization test requires a running PostgreSQL database
	// For CI/CD, set up a test database or use testcontainers
}

func TestPingDB(t *testing.T) {
	// Test with valid DB using SQLite for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open test DB: %v", err)
	}

	err = pingDB(db)
	if err != nil {
		t.Fatalf("pingDB failed with valid DB: %v", err)
	}

	// Test with closed DB (simulate invalid DB)
	rawDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get raw DB: %v", err)
	}
	rawDB.Close()

	err = pingDB(db)
	if err == nil {
		t.Fatal("pingDB should have failed with closed DB")
	}
}
