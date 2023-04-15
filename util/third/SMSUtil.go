package third

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/spf13/viper"
	"pro/middleware/log"
	"strings"
)

// 阿里云短信

func Send(phone, code string) {
	regionId := viper.GetString("sms.regionId")
	accessKeyId := viper.GetString("sms.accessKeyId")
	accessKeySecret := viper.GetString("sms.accessKeySecret")
	templateCode := viper.GetString("sms.templateCode")
	signName := viper.GetString("sms.signName")
	client, err := dysmsapi.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	// 创建API请求并设置参数
	request := dysmsapi.CreateSendSmsRequest()

	// 该参数值为假设值，请您根据实际情况进行填写
	request.PhoneNumbers = phone
	// 该参数值为假设值，请您根据实际情况进行填写
	request.SignName = signName
	request.TemplateCode = templateCode
	request.TemplateParam = "{\"code\":" + code + "}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 打印您需要的返回值，此处打印的是此次请求的 RequestId
	if !strings.EqualFold(response.Code, "OK") {
		log.Err(response.Message)
		panic("短信发送失败")
	}
}
