package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	req.Header.Add("Accept", "application/json")

	respRecorder := httptest.NewRecorder()
	r.ServeHTTP(respRecorder, req)

	if respRecorder.Code != http.StatusOK {
		t.Log(fmt.Sprintf("HTTP %v expected, but received %v", http.StatusOK, respRecorder.Code))
		t.Fail()
	}

	if respRecorder.Body == nil {
		t.Log("POST /qrcode => wasn't return body")
		t.Fail()
	}
}
