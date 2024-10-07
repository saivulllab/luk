package public

import (
	"core/db"
	"core/internal/routes/services/public/service"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	p := e.Group("/public")

	userRepo := db.GetUserRepo()
	srv := service.MakeService(userRepo)

	p.POST("/users", srv.Create)
	p.GET("/users", srv.List)
}
