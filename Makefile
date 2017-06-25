test: 
	go test ./...

deps: 
	go get -v github.com/gin-gonic/gin
	go get -v github.com/skip2/go-qrcode