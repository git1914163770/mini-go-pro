package main

import (
	"net/http"
	"web"
)

func RunServer() {
	r := web.New()
	r.GET("/", func(c *web.Context){
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})
	r.Run(":8080")
}

func main(){
	RunServer()
}
