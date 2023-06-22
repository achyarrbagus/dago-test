package routes

import (
	"backEnd/handlers"
	"backEnd/pkg/middleware"
	"backEnd/pkg/mysql"
	"backEnd/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	e.POST("/sign-in", h.Login)
	e.POST("/sign-up", h.Register)
	e.GET("/users", h.GetAllUser)
	e.GET("/user/:id", h.GetUserById)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))

}
