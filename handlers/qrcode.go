package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/skip2/go-qrcode"
)

// QrcodeHandler is a handler function to generate the QRCode
func QrcodeHandler(context *gin.Context) {
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
}

func generateQrCode(url string) (*qrcode.QRCode, error) {
	qrcode, error := qrcode.New(url, qrcode.Medium)
	if error != nil {
		return nil, error
	}

	return qrcode, error
}
