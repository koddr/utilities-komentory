.PHONY: lint security critic test

lint:
	golangci-lint run ./... --timeout 2m

security:
	gosec ./...

critic:
	gocritic check -enableAll ./...

test: lint security critic
	go test -cover ./...