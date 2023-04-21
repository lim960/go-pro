package security

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
	"pro/middleware/log"
)

// RSA SHA1WithRSA 签名、验签

func ShaSign(content string, pri *rsa.PrivateKey) (sign string) {

	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(content))
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, pri,
		crypto.SHA1, hashed)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(signature)
	return
}

func RSAVerify(origdata, ciphertext string, pub *rsa.PublicKey) (bool, error) {
	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(origdata))
	digest := h.Sum(nil)
	body, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return false, err
	}
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA1, digest, body)
	if err != nil {
		return false, err
	}
	return true, nil
}

func LoadPub(path string) (pub *rsa.PublicKey) {
	key, err := os.ReadFile(path)
	if err != nil {
		log.Err("公钥加载异常", err)
		return
	}
	block, _ := pem.Decode(key)
	if block == nil {
		log.Err("公钥加载异常,no PEM data is found")
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Err("公钥加载异常", err)
		return
	}
	pub = pubInterface.(*rsa.PublicKey)
	return
}

func LoadPri(path string) (pri *rsa.PrivateKey) {
	key, err := os.ReadFile(path)
	if err != nil {
		log.Err("公钥加载异常", err)
		return
	}
	block, _ := pem.Decode(key)
	if block == nil {
		log.Err("私钥加载异常,no PEM data is found")
		return
	}
	private, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Err("私钥加载异常", err)
		return
	}
	pri = private.(*rsa.PrivateKey)
	return
}
