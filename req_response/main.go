package main

import (
	"github.com/gkzy/gow"
	"strconv"
)

func main() {
	r := gow.Default()
	r.GET("/", Redirect)
	r.GET("/v1/user/:uid", GetUser)
	r.GET("/v1/page/user", PageUser)
	r.POST("/v1/user/:uid", SaveUser)
	r.GET("/v1/file", GetFile)
	r.GET("/v1/string", GetString)
	r.Run() // default ":8080"

}

// User demo struct
type User struct {
	UID      int64  `json:"uid"`
	Nickname string `json:"nickname"`
}

// Redirect
//		302 redirect to /v1/file
//		curl -i  http://127.0.0.1:8080
//			HTTP/1.1 302 Found
//			Content-Type: text/html; charset=utf-8
//			Location: https://gow.22v.net
//			Date: Tue, 04 Aug 2020 02:54:15 GMT
//			Content-Length: 42
//			<a href="https://gow.22v.net">Found</a>.
//
func Redirect(c *gow.Context) {
	c.Redirect(302, "https://gow.22v.net")
}

// GetUser
//	GET /v1/user/1
func GetUser(c *gow.Context) {
	idStr := c.Param("uid")
	uid, _ := strconv.ParseInt(idStr, 10, 64)
	if uid < 1 {
		c.JSON(gow.H{
			"code": 1,
			"msg":  "用户UID错误",
		})
		return
	}

	//or c.XML
	c.JSON(gow.H{
		"code": 0,
		"msg":  "success",
		"data": &User{
			UID:      uid,
			Nickname: "这是一个用户昵称",
		},
	})
}

// PageUser
//	GET http://127.0.0.1:8080/v1/page/user?page=1&limit=100
func PageUser(c *gow.Context) {
	var (
		limit int
		page  int
	)

	// c.GetString
	// c.GetInt
	// c.GetBool
	// c.GetInt64
	// c.GetFloat
	// c.GetStrings()

	limit, _ = c.GetInt("limit", 5)
	page, _ = c.GetInt("page", 1)
	if page < 1 {
		page = 1
	}

	list := make([]*User, 0)
	list = append(list, &User{
		UID:      1,
		Nickname: "用户1",
	})
	list = append(list, &User{
		UID:      2,
		Nickname: "用户2",
	})
	list = append(list, &User{
		UID:      3,
		Nickname: "用户3",
	})

	//or c.XML
	c.JSON(gow.H{
		"code": 0,
		"msg":  "success",
		"pager": gow.H{
			"page":  page,
			"limit": limit,
			"count": len(list),
		},
		"data": list,
	})
}

// SaveUser
//	POST /v1/user/1
//	request body:
//		{
//			"uid":1,
//			"nickname":"用户昵称"
//		}
func SaveUser(c *gow.Context) {
	user := new(User)
	err := c.DecodeJSONBody(&user)
	if err != nil {
		c.JSON(gow.H{
			"code": 1,
			"msg":  "获取和序列化request.body失败",
		})
		return
	}
	//or c.XML
	c.JSON(gow.H{
		"code": 0,
		"msg":  "success",
		"data": user,
	})
}

// GetFile read the server file readme.txt and output
//	GET /v1/file
func GetFile(c *gow.Context) {
	c.File("./readme.txt")
}

// GetString response string "hello gow"
func GetString(c *gow.Context) {
	c.String("hello gow:")
	// c.ServerString(404,"页面不存在")
}
