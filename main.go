package main

import (
	"github.com/danielsoro/qrcode-generate/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	startGin()
}

func startGin() {
	r := gin.New()
	r.Use(gin.Recovery())

	r.POST("/qrcode", handlers.QrcodeHandler)
	r.Run()
}
