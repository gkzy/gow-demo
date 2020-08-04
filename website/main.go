package main

import (
	"fmt"
	"github.com/gkzy/gow"
	"github.com/gkzy/gow/lib/util"
	"time"
)

func main() {
	r := gow.Default()
	r.AutoRender = true
	r.SetView("views")
	r.Static("/static", "static")

	r.Any("/", IndexPage)
	r.Any("/p/:id", OtherPage)
	r.Run()
}

// OtherPage
func OtherPage(c *gow.Context) {
	//get router param
	idStr := c.Param("id")
	//print router param
	fmt.Println(idStr)
	c.HTML("other.html")
}

// IndexPage
func IndexPage(c *gow.Context) {
	data := gow.H{}
	data["title"] = "hello gow"
	data["kewyords"] = "这是页面关键字"
	data["description"] = "这是页面介绍"
	data["date"] = util.TimeFormat(time.Now(), "YYYY-MM-DD HH:mm:ss")
	c.HTML("index.html", data)
}
