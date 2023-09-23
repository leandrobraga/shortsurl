FROM golang:1.21.1-alpine

WORKDIR /shortsurl

COPY . .

RUN go mod tidy