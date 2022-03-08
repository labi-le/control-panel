package internal

import (
	"github.com/labi-le/control-panel/structures"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDB(c *structures.Config) *DB {
	conn, err := gorm.Open(sqlite.Open(c.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &DB{db: conn}
}

// Migrate the database
func (conn *DB) Migrate() error {
	return conn.db.AutoMigrate(structures.PanelSettings{})
}

// GetSettings returns the settings
// if the settings are not found, it will save and return default settings
func (conn *DB) GetSettings() (*structures.PanelSettings, error) {
	var settings *structures.PanelSettings
	if err := conn.db.FirstOrCreate(settings).Error; err != nil {
		settings = structures.DefaultPanelSettings()
	}

	return settings, nil
}

// UpdateSettings updates the settings
func (conn *DB) UpdateSettings(settings structures.PanelSettings) error {
	return conn.db.Model(&settings).Where("_rowid_ = ?", 1).Updates(&settings).Error
}
