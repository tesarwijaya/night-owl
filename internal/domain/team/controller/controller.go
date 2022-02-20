package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tesarwijaya/night-owl/internal/domain/team/model"
	"github.com/tesarwijaya/night-owl/internal/domain/team/service"
)

type TeamController struct {
	Service service.TeamService
}

func NewTeamController(service service.TeamService) TeamController {
	return TeamController{
		Service: service,
	}
}

func (c *TeamController) SetRouter(ec *echo.Echo) {
	ec.GET("/team", c.FindAll)
	ec.GET("/team/:id", c.FindByID)
	ec.GET("/team/:id/player", c.FindTeamPlayer)
	ec.POST("/team", c.Insert)
}

func (c *TeamController) FindAll(ec echo.Context) error {
	res, err := c.Service.FindAll(ec.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, res)
}

func (c *TeamController) FindByID(ec echo.Context) error {
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

func (c *TeamController) Insert(ec echo.Context) error {
	var payload model.TeamModel

	if err := ec.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.Service.Insert(ec.Request().Context(), payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, res)
}

func (c *TeamController) FindTeamPlayer(ec echo.Context) error {
	idParam := ec.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := c.Service.FindTeamPlayer(ec.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, res)
}
