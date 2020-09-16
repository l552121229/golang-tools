package path_test

import (
	"fmt"
	"testing"

	"github.com/l552121229/golang-tools/path"
)

func TestCreate(t *testing.T) {
	fmt.Println("测试开始")
	a, err := path.Create("../tmp/pp2pp/pp")
	fmt.Println(a, err)
	a.Close()
}
