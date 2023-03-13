package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

// 上传文件
// PingExample godoc
// @Summary  上传文件接口
// @Schemes
// @Param file formData string true "表单name"
// @Description 上传文件
// @Description file 为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string} {"code":1,"msg":"\u0001个文件上传成功","url":"127.0.0.1:8080/img/12.png"}
// @Router  /user/delete      [post]
func File(c *gin.Context) {
	var filename string
	fileslice := make([]string, 10)
	form, err := c.MultipartForm()
	if err != nil {
		log.Println(err)
		return
	}
	file, err1 := os.ReadDir("./img")
	if err1 != nil {
		log.Println(err)
		return
	}
	for _, k := range file {
		fileslice = append(fileslice, k.Name())
	}
	files := form.File["file"]
	for _, file := range files {
		fmt.Println("文件名称：", file.Filename)
		for _, k := range fileslice {
			if k == file.Filename {
				c.JSON(http.StatusForbidden, gin.H{
					"msg":  "文件已经存在，请不要重复上传！",
					"code": 0,
				})
				return
			}
		}
		//上传指定目录
		c.SaveUploadedFile(file, "./img/"+file.Filename)
		filename = file.Filename

	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  string(len(files)) + "个文件上传成功",
		"code": 1,
		"url":  "127.0.0.1:8080/img/" + filename,
	})

}
