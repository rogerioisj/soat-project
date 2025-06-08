FROM golang:1.24-bookworm as builder

WORKDIR /app

copy go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/bin/api ./cmd/api/main.go

FROM debian:bookworm-20250203-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      ca-certificates tzdata libxml2-dev libxslt1-dev && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/bin/api .
COPY ./docs/swagger.yaml ./docs/

EXPOSE 8080

CMD ["./api"]