# Warhammer 🔥

> A gRPC API for Warhammer Age of Sigmar

## Intro ℹ️

> [Skip to Quick Start](https://github.com/brittonhayes/warhammer#quick-start-)

What is this?

- An API for creating, reading, deleting, and updating a database of Warhammer Age of Sigmar entities.
- A set of Protobuf schemas to generate your own gRPC client or server
- A set of OpenAPI specifications to generate your own http client or server

Why?

Most of the knowledge base for Warhammer Age of Sigmar relies on users digging through official PDF manuals. This makes it hard to efficiently find knowledge about the Warhammer universe for creators and players.

The goal of this repository is to empower the Warhammer community to make their own game companion tools in whatever language they please.

## Quick Start ✨

```shell
# Start the service with Railway
railway run

# OR start the service + database with docker-compose
docker-compose up

# OR deploy the gRPC service + HTTP REST API + PostgreSQL DB with kubernetes
kubectl apply -f deploy/warhammer.yaml
```

## Documentation 📚

The Warhammer Age of Sigmar API provides support for protocol buffers via gRPC and http via REST. View the specifications below and generate your own server or client in any language.

↔️ Protobuf - [View or download on buf.build](https://buf.build/brittonhayes/warhammer/docs)

🌐 Open API - [View Open API 3 Specification](./proto/gen/openapi/brittonhayes/warhammer/sigmar/v1)


## Deploy 🚀

Warhammer is deployed with [railway.app](https://railway.app) automatically when a change is merged into the main branch.

```shell
# Manual Deploy
railway up
```

## Development ⚙️

```shell
# Generate kubernetes manifest
# Generate go code
make build
```