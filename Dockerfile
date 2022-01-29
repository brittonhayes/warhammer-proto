# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache build-base
WORKDIR /src
COPY . .
RUN go generate

# server image
FROM golang:alpine
COPY --from=builder /src/bin/warhammer /usr/local/bin/

ARG DATABASE_READONLY_URL
ARG READONLY_USER
ARG PGDATABASE
ARG PGHOST
ARG PGPASSWORD
ARG PGPORT  
ARG PGUSER

ARG PORT
ENV PORT ${PORT}

EXPOSE 3000:3000

CMD ["/usr/local/bin/warhammer"]