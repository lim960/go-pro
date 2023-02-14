package controller

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"pro/response"
	"pro/util"
)

// Upload 上传文件
func Upload(ctx *gin.Context) {

	file, err := ctx.FormFile("file")

	if err != nil {
		panic(err)
	}
	//文件名
	fileName := util.GetFileName(16) + filepath.Ext(file.Filename)
	//保存文件
	ctx.SaveUploadedFile(file, "/"+fileName)

	response.Success(ctx, fileName)

}

// BatchUpload 批量上传文件
func BatchUpload(ctx *gin.Context) {

	form, _ := ctx.MultipartForm()
	files := form.File["file[]"]
	var list []string
	for _, file := range files {
		//文件名
		fileName := util.GetFileName(16) + filepath.Ext(file.Filename)
		//保存文件
		ctx.SaveUploadedFile(file, "/"+fileName)
		//文件名添加到切片
		list = append(list, fileName)
	}
	response.Success(ctx, list)
}
