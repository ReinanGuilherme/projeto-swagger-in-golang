package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example testando Helloworld
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /HelloWorld [get]
func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example testando HelloFriend
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} HelloFriend
// @Router /HelloFriend [get]
func HelloFriend(g *gin.Context) {
	g.JSON(http.StatusOK, "hellofriend")
}
