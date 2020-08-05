/*
实现 github.com/gkzy/gow/lib/cache_data 的 ICache接口
*/

package cache_data

import (
	"fmt"
	"github.com/gkzy/gow-demo/mem_cache/models"
)

type ProvCache struct{}

// KeyName 实现接口
func (m *ProvCache) KeyName() string {
	return "prov"
}

// PrimaryKey 实现接口
func (m *ProvCache) PrimaryKey(obj interface{}) string {
	return fmt.Sprintf("%v", obj.(*models.Prov).ID)
}

// GetAllData 取缓存中的所有数据
// 	实现接口
func (m *ProvCache) GetAllData() (data interface{}, err error) {
	return new(models.Prov).GetProvData()
}
