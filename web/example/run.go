package main

import (
	"net/http"
	"web"
)

func MyHandler(c *web.Context) {
	c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
}

func RunServer() {
	// r := web.New()
	// r.GET("/", func(c *web.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	// })
	// r.GET("/hello", MyHandler)
	// r.Run(":8080")

	r := web.New()
	r.GET("/", func(c *web.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	r.GET("/hello", func(c *web.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *web.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *web.Context) {
		c.JSON(http.StatusOK, web.Hash{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}

func main() {
	RunServer()
}
