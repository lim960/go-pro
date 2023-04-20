package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func Test1() {
	origData := []byte("123")         // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥
	fmt.Println("原文：", string(origData))

	fmt.Println("------------------ ECB模式 --------------------")
	encrypted := EncryptByECB(origData, key)
	fmt.Println("密文(base64)：", encrypted)
	decrypted := DecryptByECB(encrypted, key)
	fmt.Println("解密结果：", decrypted)

	fmt.Println("------------------ CBC模式 --------------------")
	encrypted = EncryptByCBC(origData, key)
	fmt.Println("密文(base64)：", encrypted)
	decrypted = DecryptByCBC(encrypted, key)
	fmt.Println("解密结果：", decrypted)

	fmt.Println("------------------ CFB模式 --------------------")
	encrypted = EncryptByCFB(origData, key)
	fmt.Println("密文(base64)：", encrypted)
	decrypted = DecryptByCFB(encrypted, key)
	fmt.Println("解密结果：", decrypted)
}

// EncryptByECB ECB加密
func EncryptByECB(origData, key []byte) string {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return base64.StdEncoding.EncodeToString(encrypted)
}

// DecryptByECB ECB解密
func DecryptByECB(encrypted string, key []byte) string {
	decode, _ := base64.StdEncoding.DecodeString(encrypted)
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted := make([]byte, len(decode))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(decode); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], decode[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return string(decrypted[:trim])
}

// EncryptByCBC CBC加密
func EncryptByCBC(origData, key []byte) string {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted := make([]byte, len(origData))                    // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return base64.StdEncoding.EncodeToString(encrypted)
}

// DecryptByCBC CBC解密
func DecryptByCBC(encrypted string, key []byte) string {
	decode, _ := base64.StdEncoding.DecodeString(encrypted)
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted := make([]byte, len(decode))                      // 创建数组
	blockMode.CryptBlocks(decrypted, decode)                    // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return string(decrypted)
}

// EncryptByCFB CFB加密
func EncryptByCFB(origData, key []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted := make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return base64.StdEncoding.EncodeToString(encrypted)
}

// DecryptByCFB CFB解密
func DecryptByCFB(encrypted string, key []byte) string {
	decode, _ := base64.StdEncoding.DecodeString(encrypted)
	block, _ := aes.NewCipher(key)
	if len(decode) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := decode[:aes.BlockSize]
	decode = decode[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decode, decode)
	return string(decode)
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
