.DEFAULT_GOAL := help
HEROKU_APP := twicall-0

## Setup
setup:
	go get -u -v github.com/golang/dep/cmd/dep

## Install dependencies
deps:
	dep ensure -v

## Start Server
run:
	go run *.go

## Deploy to heroku
deploy:
	heroku container:push web --app $(HEROKU_APP)

## Open website
open:
	heroku open --app $(HEROKU_APP)

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup deps run deploy help
