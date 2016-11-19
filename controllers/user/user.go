package controllerUser

import (
	"fmt"

	"github.com/labstack/echo"
)

//CreateUser creates a new user
func CreateUser(c echo.Context) error {
	u := &modelUser.User{}

	if err := c.Bind(u); err != nil {
		return c.JSON(409, err)
	}
	fmt.Println("Bind Successful")

	if err := servicesUser.CreateUser(u); err != nil {
		return c.JSON(409, err)
	}
	fmt.Println("Create User Successful")

	return c.JSON(200, u)
}
