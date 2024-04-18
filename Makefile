.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: dev
dev:
	go run ./...
