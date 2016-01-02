# -*- Makefile -*-

build: clean test package

test:
	gb test -race

package:
	gb build github.com/laher/goxc
	# See .goxc.json for parameters
	GOPATH="`pwd`:`pwd`/vendor" bin/goxc

clean:
	rm -rf bin pkg build debian
