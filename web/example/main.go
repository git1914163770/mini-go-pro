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
	r.GET("/index", func(c *web.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *web.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *web.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *web.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *web.Context) {
			c.JSON(http.StatusOK, web.Hash{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}

func main() {
	RunServer()
}
