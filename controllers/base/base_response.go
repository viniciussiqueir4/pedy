package base

import "github.com/gin-gonic/gin"

func Presenter(valid bool, messages []string, result interface{}) gin.H {
	return gin.H{
		"valid":    valid,
		"messages": messages,
		"result":   result,
	}
}
