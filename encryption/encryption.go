package encryption

import (
	"crypto/sha256"
	"fmt"
	"io"
	"math/rand"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	slatLen = 6
)

// Encryption 密码加密函数(随机盐)
// @Param password string 原密码
// Return string 加密后的函数
// Return string 加密使用的盐
func Encryption(password string) (pw string, salt string) {
	salt = RandStringBytes(slatLen)
	return EncryptionBySalt(password, salt), salt
}

// Encryption 密码加密函数(指定盐)
// @Param password string 原密码
// @Param salt string 加密使用的盐
// Return string 加密后的密码
func EncryptionBySalt(password, salt string) string {
	h := sha256.New()
	io.WriteString(h, password)
	io.WriteString(h, salt)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// RandStringBytes 生成指定长度的随机字符串
// @Param n string 生成字符串的长度
// Return string 返回生成的随机字符串
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}