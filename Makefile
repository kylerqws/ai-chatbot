.PHONY: run build migrate rollback

MAKEFLAGS += --silent

GO_ENV := CGO_ENABLED=1 GO111MODULE=on
GO_TAGS := -tags "cgo"

GO_RUN := $(GO_ENV) go run $(GO_TAGS) main.go
GO_BUILD := $(GO_ENV) go build $(GO_TAGS) -o chatbot main.go

run:
	$(GO_RUN) $(filter-out run,$(MAKECMDGOALS))

build:
	$(GO_BUILD)

migrate:
	$(GO_RUN) dev db migrate

rollback:
	$(GO_RUN) dev db rollback

%:
	@:
