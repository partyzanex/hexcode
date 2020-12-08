.PHONY: build
build:
	GO111MODULE=on CGO_ENABLED=0 go build -o hexcode main.go

.PHONY: install
install: build
	mv hexcode /usr/local/bin/hexcode
