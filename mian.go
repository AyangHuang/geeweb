package main

import "gee"

func hello(c *gee.Context) {
	urlGet := c.Query("get")
	formGet := c.PostForm("post")
	c.HTML(200, "<p>"+urlGet+"</p>"+"<p>"+formGet+"</p>")
}
func main() {
	e := gee.New()
	g1 := e.Group("/login")
	g1.Get("/get", hello)
	g2 := e.Group("/user")
	g2.POST("/post", hello)
	e.RUN(":8080")
}
