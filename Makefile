REPO_PATH = github.com/yangb8/goecsclient
DOCKER_ENV =

all: test check

deps:
	go get -v -d -t ./...
	go get -v github.com/golang/lint/golint

test: deps
	#go test -v ./... -gocheck.v

check: deps
	go vet ./...
	golint ./...

docker-test-check:
	docker run --rm \
	        -v "$$PWD":/go/src/$(REPO_PATH) \
	        -w /go/src/$(REPO_PATH) \
	        $(DOCKER_ENV) \
	        golang:1.8 \
	        make test check

.PHONY: deps install test check docker-test-check
