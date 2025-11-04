package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := pingDB(db); err != nil {
		return nil, err
	}
	return db, nil
}

func pingDB(db *gorm.DB) error {
	rawDB, err := db.DB()
	if err != nil {
		return err
	}

	if err := rawDB.Ping(); err != nil {
		return err
	}
	return nil
}
