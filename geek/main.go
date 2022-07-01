package main

import (
	"geek"
	"net/http"
)

func main() {
	r := geek.New()
	r.GET("/", func(c *geek.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *geek.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *geek.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *geek.Context) {
		c.JSON(http.StatusOK, geek.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
