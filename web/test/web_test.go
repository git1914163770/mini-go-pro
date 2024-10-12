package test

import (
	"net/http"
	"testing"
	"web"
)

func TestWeb(t *testing.T) {
	r := web.New()
	r.GET("/", func(c *web.Context){
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})
	r.Run(":8080")
}
