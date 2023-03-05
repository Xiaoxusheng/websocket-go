package test

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"testing"
)

func Test_img(t *testing.T) {
	file, err := os.Create("test.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//调用NewAlpha函数，实现image接口
	alpha := image.NewAlpha(image.Rect(0, 0, 700, 700))
	//遍历每一个像素点，设置图片
	for x := 0; x < 700; x++ {
		for y := 0; y < 700; y++ {
			alpha.Set(x, y, image.Black) //设定alpha图片的透明度

		}
	}
	//在这一行的后边，尝试调用上边提到的各种方法
	fmt.Println(alpha.At(400, 100))    //查看在指定位置的像素
	fmt.Println(alpha.Bounds())        //查看图片边界
	fmt.Println(alpha.Opaque())        //查看是否图片完全透明
	fmt.Println(alpha.PixOffset(1, 1)) //查看指定点相对于第一个点的距离
	fmt.Println(alpha.Stride)          //查看两个垂直像素之间的距离
	jpeg.Encode(file, alpha, nil)      //将alpha中的信息写入图片文件中

}
