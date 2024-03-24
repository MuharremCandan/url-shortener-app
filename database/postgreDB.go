package database

import (
	"fmt"

	"github.com/MuharremCandan/url-shortenerapp/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgre struct {
	cfg *config.Config
}

func NewPostgreDB(cfg *config.Config) IDatabase {
	return &postgre{
		cfg: cfg,
	}
}

func (d *postgre) ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", d.cfg.Database.Host, d.cfg.Database.Port, d.cfg.Database.User, d.cfg.Database.Pass,
		d.cfg.Database.Name)
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  dsn,
				PreferSimpleProtocol: true,
			},
		))
	if err != nil {
		return nil, fmt.Errorf("gorm open error: %w", err)
	}
	return db, nil
}

func (d *postgre) Ping() error {
	db, err := d.ConnectDB()
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
