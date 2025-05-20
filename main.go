package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func redirectHttp() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		target := "https://" + c.Request.Host + c.Request.URL.Path

		if c.Request.URL.RawQuery != "" {
			target += "?" + c.Request.URL.RawQuery
		}

		c.Redirect(http.StatusMovedPermanently, target)
		c.Abort()
	})

	r.Run(":80")
}

func main() {
	go redirectHttp()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Web!!!")
	})

	r.RunTLS(
		":443",
		"/certs/cert.pem",
		"/certs/privkey.pem",
	)
}
