package routers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//InitUserRoutes - Initializes all routes for user services
func InitUserRoutes(router *echo.Echo) *echo.Echo {
	//Force Signing Authorization of JWT
	var signingKey = servicesJWT.GetSigningKey()

	//Find All Users
	router.POST("/user", controllerUser.CreateUser)
	//Update User
	router.PUT("/user", controllerUser.UpdateUser, middleware.JWT([]byte(signingKey)))
	//Find Single User
	router.DELETE("/user/:username", controllerUser.DeleteUser, middleware.JWT([]byte(signingKey)))

	return router
}
