package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"pr-prueba-svc/internal/user"
	"pr-prueba-svc/internal/user/models"
	"strconv"
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

func (t *userController) GetAll(c echo.Context) error {
	list := models.User.GetAll()
	return c.JSON(http.StatusOK, list)
}
func (t *userController) GetOne(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	use, _ := models.User.GetOne(uint(id))
	return c.JSON(http.StatusOK, use)
}

func (t *userController) Update(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	u := new(user.RequestUser)
	if err := c.Bind(u); err != nil {
		return err
	}
	models.User.Update(uint(id), u.Account, u.Password)

	return c.JSON(http.StatusOK, "Registro actualizado")
}

func (t *userController) Delete(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	models.User.Delete(uint(id))
	return c.JSON(http.StatusOK, "Registro eliminado")
}
