package util

import (
	"crypto/rand"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/big"
	"strings"
	"time"
)

// Get32 32位UUID
func Get32() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

// GetRandStr 指定长度随机字符串
func GetRandStr(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	result := make([]byte, n)
	for i := range result {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		result[i] = letters[n.Int64()]
	}
	return string(result)
}

// GetFileName 生成指定长度的文件名
func GetFileName(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyz123456789")
	result := make([]byte, n)
	for i := range result {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		result[i] = letters[n.Int64()]
	}
	return string(result)
}

// GetCode 4位验证码
func GetCode() string {

	n, _ := rand.Int(rand.Reader, big.NewInt(int64(9000)))
	return fmt.Sprintf("%d", n.Int64()+1000)
}

// GetOrderNum 20位订单号  毫秒时间戳 + 6位随机数
func GetOrderNum() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(90000)))
	date := strings.ReplaceAll(time.Now().Format("20060102150405.999"), ".", "")[2:]
	return fmt.Sprintf("%s%d", date, 10000+n.Int64())
}
