.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-generate-watch
templ-generate-watch:
	templ generate --watch --proxy="http://localhost:3000" --cmd="go run ./..."

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
