FROM golang:1.21-alpine as builder
WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependedcies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# build
COPY ./ ./
#COPY app ./
RUN go build -o ./bin/app ./cmd/api/main.go
FROM cgr.dev/chainguard/glibc-dynamic

FROM alpine as runner

COPY --from=builder /usr/local/src/bin/app /
#COPY --from=builder /usr/local/src/app /
RUN mkdir config
COPY config/config-docker.yml /config/config-local.yml

CMD ["/app"]
