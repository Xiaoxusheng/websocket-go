package test

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/png"
	"log"
	"os"
	"testing"
)

func Test_img(t *testing.T) {
	width := 200
	height := 200
	canvas := image.NewRGBA(image.Rect(0, 0, width, height))

	text := "Hello, World!"
	point := fixed.Point26_6{fixed.Int26_6(10 * 64), fixed.Int26_6(50 * 64)}
	d := &font.Drawer{
		Dst:  canvas,
		Src:  image.White,
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)

	file, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, canvas)
	if err != nil {
		log.Fatal(err)
	}
}
