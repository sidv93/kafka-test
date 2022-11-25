package dto

import (
	"github.com/gin-gonic/gin"
)

type StatsCreateDTO struct {
	MemTotal     int `json:"MemTotal" binding:"required"`
	MemFree      int `json:"MemFree" binding:"required"`
	MemAvailable int `json:"MemAvailable" binding:"required"`
	SwapTotal    int `json:"SwapTotal" binding:"required"`
	SwapCached   int `json:"SwapCached" binding:"required"`
	SwapFree     int `json:"SwapFree" binding:"required"`
}

func NewStatsPost(ctx *gin.Context) (*StatsCreateDTO, error) {
	var statsDTO StatsCreateDTO

	err := ctx.ShouldBind(&statsDTO)
	if err != nil {
		return &statsDTO, err
	}

	return &statsDTO, nil
}
