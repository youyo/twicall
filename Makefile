.DEFAULT_GOAL := help
HEROKU_APP := twicall-0
VERSION := $(shell git describe --tags --abbrev=0)

## Setup
setup:
	go get -u -v github.com/golang/dep/cmd/dep
	go get -u -v  github.com/Songmu/make2help/cmd/make2help

## Install dependencies
deps:
	dep ensure -v

## Start Server
run:
	go run *.go

## Deploy to heroku
deploy:
	heroku container:push web --app $(HEROKU_APP)

## Build docker-image
build-docker-image:
	docker image build -t youyo/twicall:$(VERSION) .

## Open website
open:
	heroku open --app $(HEROKU_APP)

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: help
.SILENT:
