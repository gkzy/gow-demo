package main

import "github.com/gkzy/gow"

func main() {
	r := gow.Default()
	r.GET("/", GetUser)

	// custom error handler
	r.NoRoute(ErrorHandler)
	r.Run()
}

// GetUser
//	GET /
func GetUser(c *gow.Context) {
	c.JSON(gow.H{
		"code": 0,
		"msg":  "success",
		"data": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
	})
}

// ErrorHandler
//	other requests except GET /
func ErrorHandler(c *gow.Context) {
	c.ServerJSON(404, gow.H{
		"code": 404,
		"msg":  "request not found",
	})
}
