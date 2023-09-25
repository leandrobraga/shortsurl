FROM golang:1.21.1-alpine

WORKDIR /shortsurl

RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN apk add git 

COPY . .

RUN go mod tidy