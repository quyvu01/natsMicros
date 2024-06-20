package configurations

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"natsMicros/centralize/api/docs"
)

func SwaggerConfiguration(echo *echo.Echo) {
	docs.SwaggerInfo.Title = "Cappuccino APIs"
	docs.SwaggerInfo.Description = "Cappuccino APIs"

	echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
