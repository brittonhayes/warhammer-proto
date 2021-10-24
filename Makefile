.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: deploy
deploy:
	railway up

.PHONY: proto
proto:
	@echo "Generating code from protobuf schemas"
	buf generate --template proto/buf.gen.yaml proto

.PHONY: k8s
k8s:
	@echo "Building kubernetes manifests"
	kustomize build deploy/k8s/ -o deploy/warhammer.yaml

.PHONY: build
build: proto k8s