package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SayHello just send a hello text on /api route
func SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello")
}