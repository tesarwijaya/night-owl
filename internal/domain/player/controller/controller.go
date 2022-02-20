package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tesarwijaya/night-owl/internal/domain/player/model"
	"github.com/tesarwijaya/night-owl/internal/domain/player/service"
)

type PlayerController struct {
	Service service.PlayerService
}

func NewPlayerController(service service.PlayerService) PlayerController {
	return PlayerController{
		Service: service,
	}
}

func (c *PlayerController) SetRouter(ec *echo.Echo) {
	ec.GET("/player", c.FindAll)
	ec.GET("/player/:id", c.FindByID)
	ec.POST("/player", c.Insert)
}

func (c *PlayerController) FindAll(ec echo.Context) error {
	res, err := c.Service.FindAll(ec.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, res)
}

func (c *PlayerController) FindByID(ec echo.Context) error {
	idParam := ec.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.Service.FindByID(ec.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, res)
}

func (c *PlayerController) Insert(ec echo.Context) error {
	var payload model.PlayerModel

	if err := ec.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.Service.Insert(ec.Request().Context(), payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, res)
}
