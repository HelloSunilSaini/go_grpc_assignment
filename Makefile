#!/bin/bash

gofmt:
	go fmt ./...

test: gofmt
	go test ./... -v

run: gofmt
	docker-compose up -d --force-recreate --build app

stop:
	docker-compose down

clean:
	rm -rf client

testclient: run 
	sh scripts/client_app_test.sh
	docker-compose down
	rm -rf client
	