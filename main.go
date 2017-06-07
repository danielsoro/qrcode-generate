package main

import (
	"github.com/danielsoro/qrcode-generate/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/qrcode", handlers.QrcodeHandler)
	r.Run()
}
