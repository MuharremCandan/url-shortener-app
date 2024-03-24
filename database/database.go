package database

import "gorm.io/gorm"

type IDatabase interface {
	ConnectDB() (*gorm.DB, error)
}
