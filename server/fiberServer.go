package server

import (
	"log"
	"net"

	"github.com/MuharremCandan/url-shortenerapp/config"
	"github.com/MuharremCandan/url-shortenerapp/redirect/handlers"
	"github.com/MuharremCandan/url-shortenerapp/redirect/migrations"
	"github.com/MuharremCandan/url-shortenerapp/redirect/repository"
	"github.com/MuharremCandan/url-shortenerapp/redirect/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type server struct {
	app *fiber.App
	cfg *config.Config
	db  *gorm.DB
}

func NewFiberServer(cfg *config.Config, db *gorm.DB) IServer {
	return &server{
		app: fiber.New(),
		cfg: cfg,
		db:  db,
	}
}

// Start implements IServer.
func (s *server) Start() {

	migrate := migrations.NewMigrate(s.db)
	err := migrate.Migrate()
	if err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	repository := repository.NewRedirectRepository(s.db)
	service := service.NewRedirectService(repository)
	handler := handlers.NewHandler(service)

	s.app.Post("/", handler.Store)
	s.app.Get("/", handler.Find)

	errch := make(chan error)
	go func() {
		err := s.app.Listen(net.JoinHostPort(s.cfg.HttpServer.Host, s.cfg.HttpServer.Port))
		errch <- err
	}()

	if err := <-errch; err != nil {
		log.Fatalf("error starting server: %v", err)
	}

}
