package migrations

import (
	"github.com/MuharremCandan/url-shortenerapp/redirect/models"
	"gorm.io/gorm"
)

type IMigrate interface {
	Migrate() error
}

type migrate struct {
	db *gorm.DB
}

func NewMigrate(db *gorm.DB) IMigrate {
	return &migrate{
		db: db,
	}
}

// Migrate implements IMigrate.
func (m *migrate) Migrate() error {
	return m.db.AutoMigrate(&models.Redirect{})
}
