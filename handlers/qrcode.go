package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/skip2/go-qrcode"
)

// QRCodeRequest is a struct that represent the data request by client
type QRCodeRequest struct {
	URL string `json:"url" binding:"required"`
}

// QrcodeHandler is a handler function to generate the QRCode
func QrcodeHandler(context *gin.Context) {
	var json QRCodeRequest
	if context.BindJSON(&json) == nil {
		qrcode, error := generateQrCode(json.URL)
		if error != nil {
			context.AbortWithError(400, error)
			return
		}

		bytes, error := qrcode.PNG(256)
		if error != nil {
			context.AbortWithError(500, error)
			return
		}
		context.Render(200, render.Data{ContentType: "image/png", Data: bytes})
		return
	}
}

func generateQrCode(url string) (*qrcode.QRCode, error) {
	qrcode, error := qrcode.New(url, qrcode.Medium)
	if error != nil {
		return nil, error
	}

	return qrcode, error
}
