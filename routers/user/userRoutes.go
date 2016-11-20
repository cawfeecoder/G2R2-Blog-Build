package routers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nfrush/G2R2-Blog-Build/controllers/user"
	"github.com/nfrush/G2R2-Blog-Build/services/token"
)

//InitUserRoutes - Initializes all routes for user services
func InitUserRoutes(router *echo.Echo) *echo.Echo {
	//Force Signing Authorization of JWT
	var signingKey = servicesToken.GetSigningKey()

	//Create a user
	router.POST("/user", controllerUser.CreateUser)
	//Update a user
	router.PUT("/user", controllerUser.UpdateUser, middleware.JWT([]byte(signingKey)))
	//Delete a single user
	router.DELETE("/user/:username", controllerUser.DeleteUser, middleware.JWT([]byte(signingKey)))

	return router
}
