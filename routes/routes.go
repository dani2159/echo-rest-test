package routes

import (
	"net/http"

	"github.com/dani2159/echo-rest/controllers"
	"github.com/dani2159/echo-rest/middleware"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/list", controllers.FetchAllPegawai)

	e.GET("/pegawai", controllers.FetchAllPegawai, middleware.IsAuthenticated)
	e.POST("/pegawai", controllers.StorePegawai, middleware.IsAuthenticated)
	e.PUT("/pegawai", controllers.UpdatePegawai, middleware.IsAuthenticated)
	e.DELETE("/pegawai", controllers.DeletePegawai, middleware.IsAuthenticated)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.ChechkLogin)

	return e
}
