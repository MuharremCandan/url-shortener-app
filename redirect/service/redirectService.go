package service

import (
	"errors"
	"time"

	"github.com/MuharremCandan/url-shortenerapp/redirect/models"
	"github.com/MuharremCandan/url-shortenerapp/redirect/repository"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
)

type IRedirectService interface {
	Store(redirect *models.Redirect) (*models.Redirect, error)
	Find(code string) (*models.Redirect, error)
}

type redirectService struct {
	redirectRepository repository.IRedirectRepository
}

func NewRedirectService(redirectRepository repository.IRedirectRepository) IRedirectService {
	return &redirectService{
		redirectRepository: redirectRepository,
	}
}

// Find implements IRedirectService.
func (r *redirectService) Find(code string) (*models.Redirect, error) {
	if code == "" {
		return nil, errors.New("code cannot be empty")
	}
	return r.redirectRepository.Find(code)
}

// Store implements IRedirectService.
func (r *redirectService) Store(redirect *models.Redirect) (*models.Redirect, error) {
	if err := validate.Validate(redirect); err != nil {
		return nil, err
	}
	if redirect.URL == "" {
		return nil, errors.New("url cannot be empty")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()

	return r.redirectRepository.Store(redirect)
}
