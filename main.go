package main

import (
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func reverseProxy(target string) gin.HandlerFunc {
	url, err := url.Parse(target)
	if err != nil {
		log.Println("Reverse Proxy target url could not be parsed:", err)
		return nil
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	api := gin.Default()
	api.GET("/api/v2/sports/event/:eventId", reverseProxy("http://localhost:8080"))
	api.Run(":5000")
}
