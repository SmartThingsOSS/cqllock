# -*- Makefile -*-
VERSION := $(shell cat .goxc.json | jq -r .PackageVersion)

bootstrap:
	goxc -wc
	glide up

clean:
	rm -rf ./build

test:
	go test -race -v .

build: clean test
	goxc -bc="linux,darwin" -arch="amd64" -d build -build-ldflags "-X main.version=${VERSION}" xc

release: bump-patch build

bump-patch:
	goxc bump

bump-minor:
	goxc -dot=1 bump

bump-major:
	goxc -dot=0 bump
