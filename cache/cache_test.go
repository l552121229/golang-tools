package cache_test

import (
	"fmt"
	"testing"

	"github.com/l552121229/golang-tools/cache"
	customizeString "github.com/l552121229/golang-tools/string"
)

var cacheData *cache.Cache
var n, tp bool

func TestCreate(t *testing.T) {
	fmt.Println("测试开始")
	cacheData = cache.NewCache()
	n = true
}

func TestSet(t *testing.T) {
	if n == false {
		TestCreate(t)
	}
	cacheData.Set("你🐎逼", customizeString.Str2bytes("你🐎没了"))
	tp = true
}

func TestGet(t *testing.T) {
	if n == false {
		TestCreate(t)
	}
	if tp == false {
		TestSet(t)
	}
	data, _ := cacheData.Get("你🐎逼")

	fmt.Println(customizeString.Bytes2str(data))
}

func TestDel(t *testing.T) {
	if n == false {
		TestCreate(t)
	}
	if tp == false {
		TestSet(t)
	}

	cacheData.Del("你🐎逼")

	data, err := cacheData.Get("你🐎逼")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("删除失败", customizeString.Bytes2str(data))
	}
}
