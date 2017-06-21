package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router *MainRouter) PanelStatusHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "status.html", nil)
}

func (router *MainRouter) PanelStatusListHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "status.html", nil)
}
