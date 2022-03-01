FROM golang:1.17 AS builder

WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /web -a -ldflags '-linkmode external -extldflags "-static"' ./cmd/web

FROM scratch
COPY --from=builder /web /web

ENTRYPOINT ["/web"]
