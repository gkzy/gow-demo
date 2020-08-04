package main

import "github.com/gkzy/gow"

func main() {
	r := gow.Default()


	v1 := r.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.GET("/user/:id", GetUser)
			user.POST("/user/:id", SaveUser)
		}
	}

	v1.Use(UserAUTH())

	v2 := r.Group("/v2")
	{
		topic := v2.Group("/topic")
		{
			topic.GET("/topic/:id", GetTopic)
			topic.POST("/topic/:id", SaveTopic)
		}
	}
	v2.Use(TopicAUTH())

	r.Run()
}

//=======handler==============

func GetUser(c *gow.Context) {

}

func SaveUser(c *gow.Context) {

}

func GetTopic(c *gow.Context) {

}

func SaveTopic(c *gow.Context) {

}

// middleware

func UserAUTH() gow.HandlerFunc {
	return func(c *gow.Context) {
		//.......

		c.Next()

	}
}

func TopicAUTH() gow.HandlerFunc {
	return func(c *gow.Context) {

		//.....

		c.Next()
	}
}
