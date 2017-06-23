package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielsoro/qrcode-generate/handlers"
	"github.com/gin-gonic/gin"
)

func TestShouldReturnAByteArray(t *testing.T) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/qrcode", handlers.QrcodeHandler)

	qrcode := handlers.QRCodeRequest{URL: "https://www.example.com"}
	jsonByte, _ := json.Marshal(qrcode)

	req, _ := http.NewRequest("POST", "/qrcode", bytes.NewReader(jsonByte))
	req.Header.Add("Content-Type", "application/json")

	respRecorder := httptest.NewRecorder()
	r.ServeHTTP(respRecorder, req)

	if respRecorder.Code != http.StatusOK {
		t.Fail()
	}

	if respRecorder.Body == nil {
		t.Log("POST /qrcode not returned body")
		t.Fail()
	}
}
