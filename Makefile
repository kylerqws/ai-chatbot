.PHONY: run build migrate rollback

MAKEFLAGS += --silent
CGO_ENABLED = 1

CGO_TAGS := -tags "cgo"
GO_RUN := go run $(CGO_TAGS) main.go
GO_BUILD := go build $(CGO_TAGS) -o chatbot.exe main.go

run:
	@$(GO_RUN) $(filter-out run,$(MAKECMDGOALS))

build:
	@$(GO_BUILD)

migrate:
	@$(GO_RUN) dev db migrate

rollback:
	@$(GO_RUN) dev db rollback

%:
	@:
