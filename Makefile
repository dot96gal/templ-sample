.PHONY: dev
dev:
	go run ./...

.PHONY: templ-generate
templ-generate:
	templ generate
