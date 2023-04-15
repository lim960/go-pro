package common

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path/filepath"
	"pro/response"
	"pro/util"
	"pro/util/third"
)

// Upload 上传文件
func Upload(ctx *gin.Context) {

	file, err := ctx.FormFile("file")

	if err != nil {
		panic(err)
	}
	//图片请求头
	url := viper.GetString("oss.url")
	//文件名
	fileName := util.GetFileName(16) + filepath.Ext(file.Filename)
	//读取文件
	fileContent, _ := file.Open()
	var bytes = make([]byte, file.Size)
	fileContent.Read(bytes)
	//上传oss
	third.Upload(fileName, bytes)
	response.Success(ctx, url+fileName)

}

// BatchUpload 批量上传文件
func BatchUpload(ctx *gin.Context) {

	form, _ := ctx.MultipartForm()
	files := form.File["file[]"]
	var list []string
	//图片请求头
	url := viper.GetString("oss.url")
	for _, file := range files {
		//文件名
		fileName := util.GetFileName(16) + filepath.Ext(file.Filename)
		//读取文件
		fileContent, _ := file.Open()
		var bytes = make([]byte, file.Size)
		fileContent.Read(bytes)
		//上传oss
		third.Upload(fileName, bytes)
		//文件名添加到切片
		list = append(list, url+fileName)
	}
	response.Success(ctx, list)
}
