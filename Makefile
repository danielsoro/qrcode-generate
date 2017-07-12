test: 
	go test ./...

deps: 
	glide update

build:
	glide rebuild

install-deps:
	glide install

install:
	go install