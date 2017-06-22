package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/danielsoro/qrcode-generate/handlers"
	"github.com/gin-gonic/gin"
)

func shouldReturnAByteArray(t *testing.T) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/qrcode", handlers.QrcodeHandler)

	params := url.Values{}
	params.Add("url", "https://www.tomitribe.com")

	req, _ := http.NewRequest("POST", "/qrcode", strings.NewReader(params.Encode()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(params)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
}
