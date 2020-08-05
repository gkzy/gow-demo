package main

import (
	"fmt"
	"github.com/gkzy/gow"
	"github.com/gkzy/gow-demo/mem_cache/cache_data"
	"github.com/gkzy/gow-demo/mem_cache/models"
	"github.com/gkzy/gow/lib/cache"
	"strconv"
)

func main() {
	r := gow.Default()
	r.GET("/v1/prov", GetProv)
	r.GET("/v1/prov/:id", GetProvByID)
	r.Run()
}

// GetProv 获取所有prov数据
//	GET http://127.0.0.1:8080/v1/prov
func GetProv(c *gow.Context) {
	mc := cache.NewMemCache()
	data, err := mc.GetAll(new(cache_data.ProvCache))
	if err != nil {
		c.JSON(gow.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}

	value, ok := data.([]*models.Prov)
	if ok {
		c.JSON(gow.H{
			"code": 0,
			"msg":  "success",
			"data": value,
		})
	}
}

// GetProvByID 根据id获取单个prov信息
//	GET http://127.0.0.1:8080/v1/prov/51
func GetProvByID(c *gow.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	mc := cache.NewMemCache()
	item, err := mc.GetItem(new(cache_data.ProvCache), id)
	if err != nil {
		c.JSON(gow.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	value, ok := item.(*models.Prov)
	if ok {
		c.JSON(gow.H{
			"code": 0,
			"msg":  "success",
			"data": value,
		})
		return
	}
	c.JSON(gow.H{
		"code": 1,
		"msg":  fmt.Sprintf("没有找到 [id=%v] 的数据", id),
	})
}
