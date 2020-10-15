package string

import "unsafe"

// Str2bytes 字符串 转换为 []byte
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

// Bytes2str []byte 转换为 字符串
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
