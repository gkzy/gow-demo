package main

import (
	"github.com/gkzy/gow"
	"time"
)

func main() {
	r := gow.Default()

	//init session
	gow.InitSession()
	//use middleware
	r.Use(gow.Session())

	r.GET("/session/set", SetUser)
	r.GET("/session/get", GetUser)
	r.GET("/session/del", DelUser)

	r.GET("/cookie/set", SetTopic)
	r.GET("/cookie/get", GetTopic)
	r.GET("/cookie/del", DelTopic)

	r.Run()
}

var (
	key = "nickname"
)

func SetUser(c *gow.Context) {
	c.SetSession(key, "这是一个用户昵称")
}

func GetUser(c *gow.Context) {
	v := c.GetSessionString(key)
	c.String(v)
}

func DelUser(c *gow.Context) {
	c.DeleteSession(key)
}

//=======cookie=========

var (
	topicKey = "topic"
)

func SetTopic(c *gow.Context) {
	c.SetCookie(topicKey, "这是一个topic", int(10*time.Minute), "/", "", false, true)
}

func GetTopic(c *gow.Context) {
	v := c.GetCookie(topicKey)
	c.String(v)
}

func DelTopic(c *gow.Context) {
	c.SetCookie(topicKey, "", -1, "/", "", false, true)
}
