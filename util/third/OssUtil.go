package third

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"pro/middleware/log"
)

// 阿里云OSS

func Upload(fileName string, content []byte) {
	endpoint := viper.GetString("oss.endpoint")
	accessKeyID := viper.GetString("oss.accessKeyID")
	accessKeySecret := viper.GetString("oss.accessKeySecret")
	bucketName := viper.GetString("oss.bucketName")
	// 创建OSSClient实例。
	client, _ := oss.New(endpoint, accessKeyID, accessKeySecret)

	// 填写存储空间名称，例如examplebucket。
	bucket, _ := client.Bucket(bucketName)

	// 上传
	err := bucket.PutObject("lego/"+fileName, bytes.NewReader(content))
	if err != nil {
		log.Err("文件上传失败", err)
		panic("文件上传失败")
	}
}
