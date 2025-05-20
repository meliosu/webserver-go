package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Web!!!")
	})

	r.RunTLS(
		":443",
		"/etc/letsencrypt/live/meliosu.ru/fullchain.pem",
		"/etc/letsencrypt/live/meliosu.ru/privkey.pem",
	)
}
