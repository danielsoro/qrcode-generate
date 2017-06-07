package main

import (
	"github.com/skip2/go-qrcode"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/gin-gonic/gin.v1/render"
)

func main() {
	r := gin.Default()

	r.POST("/qrcode", func(context *gin.Context) {
		url := context.PostForm("url")
		qrcode, error := generateQrCode(url)
		if error != nil {
			panic(error)
		}

		bytes, error := qrcode.PNG(256)
		if error != nil {
			panic(error)
		}

		context.Render(200, render.Data{ContentType: "image/png", Data: bytes})
	})

	r.Run()
}

func generateQrCode(url string) (*qrcode.QRCode, error) {
	qrcode, error := qrcode.New(url, qrcode.Medium)
	if error != nil {
		return nil, error
	}

	return qrcode, error
}
