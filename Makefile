
HAS_GLIDE := $(shell command -v glide;)

.PHONY: test
test:
	go test $(go list ./... | grep -v /vendor/)

.PHONY: setup
setup:
ifndef HAS_GLIDE
	go get -u github.com/Masterminds/glide
endif
	glide up
	glide install