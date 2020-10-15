package string_test

import (
	"fmt"
	"testing"

	customizeString "github.com/l552121229/golang-tools/string"
)

func TestStr2bytes(t *testing.T) {
	fmt.Println("开始测试")
	bt := customizeString.Str2bytes("test string")
	fmt.Println(bt)
}

func TestBytes2str(t *testing.T) {
	fmt.Println("开始测试")
	bt := customizeString.Bytes2str([]byte{116, 101, 115, 116, 32, 115, 116, 114, 105, 110, 103})
	fmt.Println(bt)
}
