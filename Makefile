.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: build
build:
	docker build . -t brittonhayes/warhammer:latest

.PHONY: deploy
deploy:
	railway up

.PHONY: generate
generate:
	cd proto && buf generate