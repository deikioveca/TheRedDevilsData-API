FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o theRedDevilsData-cli ./cli && \
    go build -o theRedDevilsData-web ./api

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/theRedDevilsData-cli .
COPY --from=builder /app/theRedDevilsData-web .
COPY .env .

ENTRYPOINT ["/bin/sh"]
