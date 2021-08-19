package database

import (
	"github.com/glauber-silva/farmers-markets/internal/markets"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates the database and creates the market table
func MigrateDB(db *gorm.DB) error {
	if r := db.AutoMigrate(&markets.Market{}); r.Error != nil {
		return r.Error
	}
	return nil
}