package routes

import (
	"github.com/labstack/echo"
	"pr-prueba-svc/internal/user/controllers"
)

func User(e *echo.Echo) {
	r := e.Group("/user")
	/*middlewares*/
	r.POST("/add", controllers.User.Add)
	r.PUT("/update/:id", controllers.User.Update)
	r.DELETE("/del/:id", controllers.User.Delete)
	r.GET("/getone/:id", controllers.User.GetOne)
	r.GET("/getall", controllers.User.GetAll)
}
