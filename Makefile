.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test ./...

.PHONY: dev
dev: templ-generate
	go run ./...

.PHONY: air
air:
	air -c ./.air.toml
