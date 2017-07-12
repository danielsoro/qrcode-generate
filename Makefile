
HAS_GLIDE := $(shell command -v glide;)
SRC = $(shell go list ./... | grep -v "/vendor")

.PHONY: test
test:
	go test $(SRC)

.PHONY: setup
setup:
ifndef HAS_GLIDE
	go get -u github.com/Masterminds/glide
endif
	glide up
	glide install
