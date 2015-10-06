# -*- Makefile -*-

version=0.0.1
deb_file=cqllock_${version}_amd64.deb

build: test package

test:
	go test

package:
	goxc -bc="linux" -arch="amd64" -d build -pv="${version}"

clean:
	rm -rf build
