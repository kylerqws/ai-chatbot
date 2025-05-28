.PHONY: run build

MAKEFLAGS += --silent
CGO_ENABLED = 1

GO_RUN := ./chatbot.exe
GO_BUILD := go build -tags "cgo" -o chatbot.exe main.go

run: build
	- $(GO_RUN) $(filter-out run,$(MAKECMDGOALS))

build:
	- $(GO_BUILD)

%:
	@:
