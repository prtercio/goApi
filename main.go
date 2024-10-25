package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/test/:id", func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		if id != "123" {
			err, code := createErrorRequest()
			c.JSON(code, err)
			return
		}
		err, code := createOKRequest()
		c.JSON(code, err)

	})
	r.Run(":8081")
}

func createErrorRequest() (gin.H, int) {
	return gin.H{
		"message": "error request",
	}, http.StatusOK
}

func createOKRequest() (gin.H, int) {
	return gin.H{
		"message": "request OK",
	}, http.StatusOK
}
