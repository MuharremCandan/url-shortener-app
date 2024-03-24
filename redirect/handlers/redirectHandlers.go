package handlers

import (
	"net/http"

	"github.com/MuharremCandan/url-shortenerapp/redirect/entities"
	"github.com/MuharremCandan/url-shortenerapp/redirect/models"
	"github.com/MuharremCandan/url-shortenerapp/redirect/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IHandler interface {
	Find(ctx *fiber.Ctx) error
	Store(ctx *fiber.Ctx) error
}

type handler struct {
	redirectService service.IRedirectService
}

func NewHandler(redirectService service.IRedirectService) IHandler {
	return &handler{
		redirectService: redirectService,
	}
}

// Find implements IHandler.
func (h *handler) Find(ctx *fiber.Ctx) error {
	var getRedirectReqDTO entities.GetRedirectRequestDTO
	if err := ctx.BodyParser(&getRedirectReqDTO); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	redirect, err := h.redirectService.Find(getRedirectReqDTO.Code)
	if err != nil {
		if err.Error() == "code cannot be empty" {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"redirect": entities.GetRedirectResponseDTO{
			Code: redirect.Code,
			URL:  redirect.URL,
		},
	})
}

// Store implements IHandler.
func (h *handler) Store(ctx *fiber.Ctx) error {
	createRedirectDTO := entities.CreateRedirectDTO{}
	if err := ctx.BodyParser(&createRedirectDTO); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	redirect, err := h.redirectService.Store(&models.Redirect{
		URL: createRedirectDTO.URL,
	})
	if err != nil {
		if err.Error() == "url cannot be empty" {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"redirect": entities.CreatedRedirectResponseDTO{
			Code: redirect.Code,
			URL:  redirect.URL,
		},
	})
}
