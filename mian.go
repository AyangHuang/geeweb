package main

import "gee"

func hello(c *gee.Context) {
	urlGet := c.Query("get")
	formGet := c.PostForm("post")
	c.HTML(200, "<p>"+urlGet+"</p>"+"<p>"+formGet+"</p>")
}
func main() {
	e := gee.New()
	e.POST("/web", hello)
	e.RUN(":8080")
}
