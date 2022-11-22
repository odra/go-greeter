SHELL := /bin/bash

build/plugins:
	@go build -buildmode=plugin -o build/plugin_en.so contrib/plugins/plugin_en/plugin_en.go 
	@go build -buildmode=plugin -o build/plugin_ptbr.so contrib/plugins/plugin_ptbr/plugin_ptbr.go

build/bin:
	@go build -o build/greet bin/greet.go

build/clean:
	@rm -rf ./build/

build: build/plugins build/bin

.PHONY: build/plugins build/bin build/clean build
