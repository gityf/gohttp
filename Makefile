GOPATH:=$(CURDIR)
export GOPATH

all: output

fmt:
	#gofmt -l -w -s src/

dep:fmt

build:dep
	go build -o bin/simpleserver main

clean:
	rm -rf output
	rm -rf bin/simpleserver
output:build
	mkdir -p output/bin
	mkdir -p output/conf
	mkdir -p output/log
	mkdir -p output/web
	cp -r bin/simpleserver output/bin/
	cp -r conf/* output/conf/
	cp -r web/* output/web/

