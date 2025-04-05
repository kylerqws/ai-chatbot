.PHONY: run build migrate

MAKEFLAGS += --silent
CGO_ENABLED = 1

run:
	@go run -tags "cgo" main.go $(filter-out run,$(MAKECMDGOALS))

build:
	@go build -tags "cgo" -o chatbot.exe main.go

migrate:
	@if not "$(findstring up, $(MAKECMDGOALS))" == "" ( \
		$(MAKE) run dev db migrate \
	) else if not "$(findstring down, $(MAKECMDGOALS))" == "" ( \
		$(MAKE) run dev db rollback \
	)

%:
	@:
