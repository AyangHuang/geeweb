package main

import (
	"gee"
)

func hello(c *gee.Context) {
	urlGet := c.Query("get")
	formGet := c.PostForm("post")
	c.HTML(200, "<p>"+urlGet+"</p>"+"<p>"+formGet+"</p>")
}
func main() {
	e := gee.New()
	if g1 := e.Group("/login1"); g1 != nil {
		g1.Use(gee.Logger())
		g1.Get("/get", hello)
		if g2 := g1.Group("/login2"); g2 != nil {
			g2.Use(gee.Logger())
			g2.POST("/post", hello)
		}
	}
	e.POST("/post", hello)
	e.RUN(":8080")

}
