package main

import (
	gee2 "github.com/dvecccc/learn/GeeWeb/gee"
	"net/http"
)

func main() {
	r := gee2.New()
	r.Get("/", func(c *gee2.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World!</h1>")
	})
	r.Get("/hello", func(c *gee2.Context) {
		c.String(http.StatusOK, "Hello %s, you are at %s", c.Query("name"), c.Path)
	})
	r.Post("/login", func(c *gee2.Context) {
		c.JSON(http.StatusOK, gee2.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	_ = r.Run(":8082")
}
