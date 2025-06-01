.PHONY: run build migrate rollback

MAKEFLAGS += --silent
CGO_ENABLED = 1

GO_TAGS := -tags "cgo"
GO_MAIN_FILE := main.go
GO_BIN_FILE := chatbot
GO_RUN := go run $(GO_TAGS) $(GO_MAIN_FILE)
GO_BUILD := go build $(GO_TAGS) -o $(GO_BIN_FILE) $(GO_MAIN_FILE)

run:
	- $(GO_RUN) $(filter-out run,$(MAKECMDGOALS))

build:
	- $(GO_BUILD)

migrate:
	- $(GO_RUN) dev db migrate

rollback:
	- $(GO_RUN) dev db rollback
