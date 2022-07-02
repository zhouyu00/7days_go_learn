package main

import (
	"fmt"
	"geek"
	"html/template"
	"log"
	"net/http"
	"time"
)

func onlyForV2() geek.HandlerFunc {
	return func(c *geek.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s for group v2 %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := geek.New()
	r.Use(geek.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "geekktutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *geek.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *geek.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", geek.H{
			"title":  "geek",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *geek.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", geek.H{
			"title": "geek",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
