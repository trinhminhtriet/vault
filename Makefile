NAME    := vault
PACKAGE := github.com/trinhminhtriet/$(NAME)
DATE    :=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT     := $(shell [ -d .git ] && git rev-parse --short HEAD)
VERSION := $(shell git describe --tags)

default: build

tidy:
	go get -u && go mod tidy

build:
	CGO_ENABLED=0 go build \
	-ldflags "-s -w -X '${PACKAGE}/internal.VERSION=${VERSION}' -X '${PACKAGE}/internal.DATE=${DATE}'" \
	-a -tags netgo -o dist/${NAME} main.go

release:
	goreleaser build --clean --snapshot --single-target

release-all:
	goreleaser build --clean --snapshot

link:
	ln -sf ${PWD}/dist/${NAME} /usr/local/bin/${NAME}
	which ${NAME}

clean:
	$(RM) -rf dist

.PHONY: default tidy build release link
