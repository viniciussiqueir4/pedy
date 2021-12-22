package controllers

import "github.com/gin-gonic/gin"

func NewUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
