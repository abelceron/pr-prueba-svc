package bootstrap

import (
	"github.com/labstack/echo"
	config "github.com/spf13/viper"
	"pr-prueba-svc/internal/user/routes"
	"pr-prueba-svc/kit/libs"
)

func init() {
	config.AddConfigPath("./kit/config-mysql")
	dbConfig := libs.Configure("./kit/config-mysql", "mysql")
	libs.DB = dbConfig.InitMysqlDB()
}
func mainMiddlewares(e *echo.Echo) {

}
func mainRoutes(e *echo.Echo) {
	routes.User(e)
}
func Run() {
	e := echo.New()
	mainMiddlewares(e)
	mainRoutes(e)
	e.Logger.Fatal(e.Start(":80"))
}
