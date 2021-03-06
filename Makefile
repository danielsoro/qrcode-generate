
HAS_GLIDE := $(shell command -v glide;)
.PHONY: test
test:
	go test ./...

.PHONY: setup
setup:
ifndef HAS_GLIDE
	go get -u github.com/Masterminds/glide
endif
	glide up
	glide install
