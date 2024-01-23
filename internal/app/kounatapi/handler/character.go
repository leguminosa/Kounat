package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/leguminosa/kounat/internal/module"
	"github.com/leguminosa/kounat/internal/tools/convert"
	echoHelper "github.com/leguminosa/kounat/internal/tools/ecxo/helper"
)

type CharacterHandler struct {
	module module.CharacterModule
}

func NewCharacter(module module.CharacterModule) *CharacterHandler {
	return &CharacterHandler{
		module: module,
	}
}

func (h *CharacterHandler) GetByID(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		id  = convert.ToInt(c.Param("id"))
	)

	result, err := h.module.GetByID(ctx, id)
	if err != nil {
		c.Logger().Error(err)
		return echoHelper.InternalServerError(c, "something went wrong")
	}

	return echoHelper.OK(c, result)
}
