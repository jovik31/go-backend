FROM golang:alpine AS builder

WORKDIR /go-backend
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/backend

FROM alpine:latest

WORKDIR /
COPY --from=builder /go-backend/bin/* /
