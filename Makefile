default: test

get-deps:
	go get github.com/mvillalba/go-openexchangerates/oxr

test:
	go test ./... -v

package: build
	rm -rf ./dist/release/*

	$(eval VERSION := $(shell ./dist/binaries/darwin/oxr --version))

	zip -j ./dist/release/oxr_darwin_$(VERSION).zip ./dist/binaries/darwin/*
	zip -j ./dist/release/oxr_linux_$(VERSION).zip ./dist/binaries/linux/*
	zip -j ./dist/release/oxr_windows_$(VERSION).zip ./dist/binaries/windows/*

build:
	make get-deps
	make test

	rm -rf ./dist/binaries/*
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o ./dist/binaries/darwin/oxr .
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./dist/binaries/linux/oxr .
	CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -o ./dist/binaries/windows/oxr.exe .
