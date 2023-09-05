GIT_VERSION?=$(shell git describe --tags --always --abbrev=42 --dirty)

binary: bin
	go build -ldflags "-X gitlhub.com/athoune/docker-dynasty-go/version.version=$(GIT_VERSION)" \
		-o bin/dynasty

bin:
	mkdir -p bin

linux:
	make binary GOOS=linux

upx:
	upx bin/dynasty

test:
	go test -v gitlhub.com/athoune/docker-dynasty-go/dynasty
