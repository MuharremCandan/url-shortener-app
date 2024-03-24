package repository

import (
	"github.com/MuharremCandan/url-shortenerapp/redirect/models"
	"gorm.io/gorm"
)

type IRedirectRepository interface {
	Store(redirect *models.Redirect) (*models.Redirect, error)
	Find(code string) (*models.Redirect, error)
}

type redirectRepository struct {
	db *gorm.DB
}

func NewRedirectRepository(db *gorm.DB) IRedirectRepository {
	return &redirectRepository{
		db: db,
	}
}

// Find implements IRedirectRepository.
func (r *redirectRepository) Find(code string) (*models.Redirect, error) {
	var redirect models.Redirect
	err := r.db.Where("code =?", code).First(&redirect).Error
	return &redirect, err
}

// Store implements IRedirectRepository.
func (r *redirectRepository) Store(redirect *models.Redirect) (*models.Redirect, error) {
	err := r.db.Create(redirect).Error
	return redirect, err
}
