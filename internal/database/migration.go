package database

import (
	"github.com/glauber-silva/farmers-markets/internal/markets"
	"gorm.io/gorm"
)

// MigrateDB - migrates the database and creates the market table
func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&markets.Market{}); err != nil {
		return err
	}
	return nil
}