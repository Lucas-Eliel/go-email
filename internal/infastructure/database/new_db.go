package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=email_dev password=d4#rt6 dbname=email_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Fail to open database")
	}

	return db
}
