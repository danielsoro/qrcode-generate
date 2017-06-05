package main

import "github.com/skip2/go-qrcode"

func main() {
	qrCode, error := qrcode.New("http://www.twitter.com/dvlc_", qrcode.Medium)
	if error != nil {
		panic(error)
	}

	qrCode.WriteFile(250, "test.png")
}
