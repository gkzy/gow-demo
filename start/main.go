package main

import (
	"github.com/gkzy/gow"
)

func main() {
	r := gow.Default()

	r.GET("/", func(c *gow.Context) {
		c.JSON(gow.H{
			"code": 0,
			"msg":  "success",
		})
	})

	r.GET("/hello", Hello)

	r.Run()
}

func Hello(c *gow.Context) {
	c.String("hello...")
}
