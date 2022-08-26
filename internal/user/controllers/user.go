package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"pr-prueba-svc/internal/user"
	"pr-prueba-svc/internal/user/models"
)

var User *userController

type userController struct {
}

func (t *userController) Add(c echo.Context) error {
	u := new(user.RequestUser)
	if err := c.Bind(u); err != nil {
		return err
	}
	models.User.Add(u.Account, u.Password)
	return c.JSON(http.StatusOK, "Registro Creado")
}
