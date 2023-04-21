package security

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/ZZMarquis/gm/sm2"
	"strings"
)

// SM2withSM3 签名、验签

func Test() {
	//生成密钥对
	hexPri, basePub := Generate()
	println("******* 生成 hex私钥：" + hexPri)
	println("******* 生成 base公钥：" + basePub)
	//Hex私钥转私钥对象
	pri := HexToPri(hexPri)
	//base64公钥转公钥对象
	pub := Base64ToPub(basePub)
	//私钥生成公钥
	base64Pub := PriToBase64Pub(pri)
	println("******* 私钥生成base64公钥：" + base64Pub)
	//私钥转hex
	priHex := PriToHex(pri)
	println("******* 私钥转hex：" + priHex)
	//公钥转base64
	pubBase := PubToBase64(pub)
	println("******* 公钥转base64：" + pubBase)
	//签名
	sign := Sign("abc=123", pri)
	println("******* 签名：" + sign)
	//验签
	ver := Verify("abc=123", sign, pub)
	println(fmt.Sprintf("******* 验签：%t", ver))
}

// Sign 签名
func Sign(data string, pri *sm2.PrivateKey) string {
	signature, err := sm2.Sign(pri, []byte("1234567812345678"), []byte(data))
	if err != nil {
		panic("云闪付签名错误")
	}
	// 转 base64
	sign := base64.StdEncoding.EncodeToString(signature)
	return sign
}

// Verify 验签
func Verify(data, sign string, pub *sm2.PublicKey) bool {
	sign1, _ := base64.StdEncoding.DecodeString(sign)
	return sm2.Verify(pub, []byte("1234567812345678"), []byte(data), sign1)
}

// HexToPri Hex私钥转私钥对象
func HexToPri(priStr string) *sm2.PrivateKey {
	// 解码hex私钥
	privateKeyByte, _ := hex.DecodeString(priStr)
	// 转成go版的私钥
	pri, err := sm2.RawBytesToPrivateKey(privateKeyByte)
	if err != nil {
		panic("云闪付私钥加载异常")
	}
	return pri
}

// Base64ToPub base64公钥转公钥对象
func Base64ToPub(pubStr string) *sm2.PublicKey {
	decode, _ := base64.StdEncoding.DecodeString(pubStr)
	pubHex := hex.EncodeToString(decode)
	pubHex = strings.ReplaceAll(pubHex, "3059301306072a8648ce3d020106082a811ccf5501822d03420004", "")
	pubByte, _ := hex.DecodeString(pubHex)
	pub, _ := sm2.RawBytesToPublicKey(pubByte)
	return pub
}

// PriToHex 私钥转hex
func PriToHex(pri *sm2.PrivateKey) string {
	hex := hex.EncodeToString(pri.GetRawBytes())
	return hex
}

// PubToBase64 公钥转base64
func PubToBase64(pub *sm2.PublicKey) string {
	pub.GetRawBytes()
	bytes := pub.X.Bytes()
	bytes = append(bytes, pub.Y.Bytes()...)
	pubHex := "3059301306072a8648ce3d020106082a811ccf5501822d03420004" + hex.EncodeToString(bytes)
	decode, _ := hex.DecodeString(pubHex)
	base64Pub := base64.StdEncoding.EncodeToString(decode)
	return base64Pub
}

// PriToBase64Pub 私钥生成公钥
func PriToBase64Pub(pri *sm2.PrivateKey) string {
	pub := sm2.CalculatePubKey(pri)
	base64Pub := PubToBase64(pub)
	return base64Pub
}

// Generate 生成密钥对 hex私钥 base64公钥
func Generate() (string, string) {
	pri, pub, _ := sm2.GenerateKey(rand.Reader)
	return PriToHex(pri), PubToBase64(pub)
}
