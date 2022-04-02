package controller

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthzController struct {
	Db *sql.DB
}

func NewHealthzController(db *sql.DB) HealthzController {
	return HealthzController{
		Db: db,
	}
}

func (c *HealthzController) SetRouter(ec *echo.Echo) {
	ec.GET("/healthz", c.Healthz)
}

func (c *HealthzController) Healthz(ec echo.Context) error {
	DBStatus := "UP!"
	err := c.Db.Ping()
	if err != nil {
		DBStatus = err.Error()
	}

	return ec.JSON(http.StatusOK, map[string]string{
		"DBStatus": DBStatus,
	})
}
