package security

import (
	"github.com/golang-module/dongle"
)

// 私钥加密
func EncryptByPriKey(data string) string {

	return dongle.Encrypt.FromString(data).ByRsa(pri).ToBase64String()
}

// 公钥解密
func DecryptByPubKey(data string) string {
	return dongle.Decrypt.FromBase64String(data).ByRsa(pub).ToString()
}

// 公钥加密
func EncryptByPubKey(data string) string {
	return dongle.Encrypt.FromString(data).ByRsa(pub).ToBase64String()
}

// 私钥解密
func DecryptByPriKey(data string) string {
	return dongle.Decrypt.FromBase64String(data).ByRsa(pri).ToString()
}

func catch() {
	if err := recover(); err != nil {
		println(err)
	}
}

var pri = `
-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAJlDYH0MoSaHRw04loTi2bl2Jaoz
O0I4ChaoWhuGiJhnx6pRmQW5ojT3pWaE+UA5VEW0tQEzuqawLr6zTtp36x9NeIolvQkZHLudjo8G
bkLQuA+HVUQ43GL6+eepSA/mfUIA6brq7rDWP7fDe8SWua9s4V6tR/AtGCs5TvEYCCX7AgMBAAEC
gYBVFZmYco14VTt1tIejaEjU9Ck+zshEH9ZB895qT4q/iUXIYRphmkfZve399y5koC8Pr52Y+D3T
0hVxWxwYnuBRFLlMnUlyveonGD3bncI0YFvC0eHzwnWagOGsvdDD+cCCT0a6/0+iieF5jrryPYsY
/mP/chMFSpckMeBxpRGBsQJBAPP2LUQf3ONeGAC+oL3ZibWnrEWV3SydZnzXvfwmqMqRFCDksPzD
OE3GfDEvZzcB08l69zt24JFnUYI/dwYp/ikCQQCg03XVs33SJbL0sgN+VQerJsTPWMnQFWGWD/9y
wt158AVDadv9iFROjRCXtKRokUKlRY9cWlCIzzwIUx6k5z+DAkEA3VvM3Nhwa5mv+9T8MucU7c/T
H1yIz/eNy89R4l4Nn6ed9O6srNxR1Tg47cQOSjoNOe6qL7mAsE5oBd+iFuS5aQJAaVn8X9AjxOzD
LP4Lwc8LpfdQh49fLHtFINs7+D5kfQNZP07yOEP9DjPkQayo4oL9iGxnvBTBms0+QynH8jg15wJB
ALOvJmUw5KEh7zrGE6Rp8/1tnv2ZHPeUHS07n4Tas68vBqA7i3MJ9oYDO6ol7yH09OxNxt72KuaP
mD/AcOw/sJM=
-----END PRIVATE KEY-----`

var pub = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCZQ2B9DKEmh0cNOJaE4tm5diWqMztCOAoWqFob
hoiYZ8eqUZkFuaI096VmhPlAOVRFtLUBM7qmsC6+s07ad+sfTXiKJb0JGRy7nY6PBm5C0LgPh1VE
ONxi+vnnqUgP5n1CAOm66u6w1j+3w3vElrmvbOFerUfwLRgrOU7xGAgl+wIDAQAB
-----END PUBLIC KEY-----`
