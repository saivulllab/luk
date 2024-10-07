package service

import (
	"core/db/repository/user"
	"core/internal/routes/services/public/service/request"
	"github.com/labstack/echo/v4"
	"net/http"
	"validator"
)

type Service struct {
	repo user.Repository
}

func MakeService(repo user.Repository) *Service {
	if repo == nil {
		panic("user repo is nil")
	}

	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(e echo.Context) error {
	req := new(request.User)

	if err := e.Bind(req); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := validator.NewCustomValidator().ValidateWithContext(req, e); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	res := req.ToEntity()

	if err := s.repo.Create(e.Request().Context(), res); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return e.JSON(http.StatusCreated, res)
}

func (s *Service) List(e echo.Context) error {
	res, err := s.repo.List(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return e.JSON(http.StatusOK, res)
}
