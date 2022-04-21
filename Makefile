.DEFAULT_GOAL := help

OutDir=bin/kong-cli


BuildCmd=go build -a -o $(OutDir) -gcflags=all="-l" -ldflags="-w -s"

.PHONY: build
build:
	mkdir -p bin
	$(BuildCmd) .

.PHONY: run
run:
	go run main.go

.PHONY: help
help:
	@(                                     \
	  echo "TARGETS";                      \
	  echo "-------";                      \
	  echo "build,builds cli"; \
	  echo "run,Runs the cli";                                   \
	) | column -s, -t