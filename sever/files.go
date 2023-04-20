package server

import (
	"Gin/models"
	"Gin/utility"
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
// @Param file formData file true "表单name"
// @Param token header string true "token"
// @Description 上传文件
// @Description file token为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{"code":1,"msg":"\u0001个文件上传成功","url":"127.0.0.1:8080/img/12.png"}
// @Router  /user/file      [post]
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
				c.JSON(http.StatusOK, gin.H{
					"msg":  "文件已经存在，请不要重复上传！",
					"code": 0,
					"url":  "http://127.0.0.1:8080/img/" + filename,
				})
				return
			}
		}
		//index := strings.LastIndex(filename, ".")
		//if index != -1 {
		//	extension := filename[index+1:]
		//	fmt.Println(extension)
		//}
		//上传指定目录
		c.SaveUploadedFile(file, "./img/"+file.Filename)
		log.Println(file.Filename)
		filename = file.Filename

	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  string(len(files)) + "个文件上传成功",
		"code": 200,
		"url":  "http://127.0.0.1:8080/img/" + filename,
	})

}

// SetHeadPicture
// PingExample godoc
// @Summary 上传头像接口
// @Tags 公共方法
// @Param HeadPicture formData file true "表单name"
// @Param token header string true "token"
// @Description file token为必填
// @Accept multipart/form-data
// @Produce json
// @Success 200 { string } "{"account": "3169387148", "code": 200, "msg": "注册成功"}"
// @Router  /user/setheadpicture      [post]
func SetHeadPicture(c *gin.Context) {
	use := c.MustGet("use")
	user := use.(*utility.User)

	// 单文件
	file, _ := c.FormFile("HeadPicture")
	log.Println(file.Filename)
	filelist, err1 := os.ReadDir("./img")
	if err1 != nil {
		log.Println(err1)
		return
	}
	for _, k := range filelist {
		if file.Filename == k.Name() {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "不能使用相同头像!",
				"code": 1,
			})
		}
		break
	}
	//dst := "./" + file.Filename
	//// 上传文件至指定的完整文件路径
	//c.SaveUploadedFile(file, dst)

	c.SaveUploadedFile(file, "./img/"+file.Filename)
	filename := file.Filename
	url := "http://127.0.0.1:8080/img/" + filename

	err := models.SetHeadPicture(url, (user).Indently)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "头像设置成功！",
		"code": 200,
		"url":  url,
	})

}
