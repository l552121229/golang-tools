package path

import (
	"fmt"
	"os"
	"path"
)

// Create file
// 返回创建好的文件
func Create(fileName string) (*(os.File), error) {
	f, err := os.Create(fileName)

	if err != nil {
		_, StatusErr := os.Stat(fileName)
		if StatusErr == nil {
			return nil, err
		}
		if os.IsNotExist(StatusErr) {
			err = MkdirAll(path.Dir(fileName), 0766)
			fmt.Println(err)
			if err != nil {
				return nil, err
			}
		}
		return os.Create(fileName)
	}

	return f, nil
}

// Mkdir 相当于os.Mkdir()
func Mkdir(dir string, perm os.FileMode) error {
	return os.Mkdir(dir, perm)
}

// MkdirAll 创建文件夹 相当于os.MkdirAll()
func MkdirAll(dir string, perm os.FileMode) error {
	fDir := path.Dir(dir)

	if fDir != "." && fDir != ".." && fDir != "/" {
		err := MkdirAll(fDir, perm)
		if err != nil {
			return err
		}
	}
	_, err := os.Stat(dir)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		fmt.Println(dir)
		err = Mkdir(dir, perm)
		if err != nil {
			return err
		}
	}
	return nil
}
