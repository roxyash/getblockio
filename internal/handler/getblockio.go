package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/roxyash/getblock/pkg/getblock"
	"github.com/roxyash/getblock/internal/helpers"
	"github.com/sirupsen/logrus"
)

// @Summary getblockio
// @Tags getblockio
// @Description getblockio
// @ID getblockio
// @Accept  json
// @Produce  json
// @Param        apikey_id   path      string  true  "Apikey Id"
// @Success 200 {object} types.ResponseCheckBalanceInfo
// @Failure 400,404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Failure default {object} types.ErrorResponse
// @Router /getblockio/{apikey_id} [post]
func (h *Handler) checkBalance(c *gin.Context) {
	apikeyId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid apikey id param")
		return
	}
	blockNumber, err := getblock.GetBlockNumber(apikeyId)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	blockNumberInfo, err := getblock.GetBlockInfoByNumber(apikeyId, blockNumber.Result, true)

	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	balanceInfo, err := getblock.GetBalanceByMiner(apikeyId, blockNumberInfo.Result.Miner)

	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Что-то делаем, чтобы посчитать разницу между блоками. Чтобы узнать изменился ли баланс или нет. . .

	logrus.Infof("%v", balanceInfo.Result)

	c.JSON(http.StatusOK, map[string]string{
		"Address": blockNumberInfo.Result.Miner,
	})

}
