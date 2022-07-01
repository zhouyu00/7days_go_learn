package main

import (
	"geek"
	"net/http"
)

func main() {
	r := geek.New()
	r.GET("/index", func(c *geek.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *geek.Context) {
			c.HTML(http.StatusOK, "<h1>Hello geek</h1>")
		})

		v1.GET("/hello", func(c *geek.Context) {
			// expect /hello?name=geekktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *geek.Context) {
			// expect /hello/geekktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *geek.Context) {
			c.JSON(http.StatusOK, geek.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
