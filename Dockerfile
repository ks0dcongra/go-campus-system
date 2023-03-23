FROM golang:1.18.0-alpine3.15 as builder

LABEL maintainer="Shehomebow"

RUN apk update && apk add --no-cache git

RUN mkdir -p /app

WORKDIR /app

COPY ./go.mod ./go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

COPY . .

RUN go build -o andy_training

ENTRYPOINT  ["/app/andy_training"]

EXPOSE 9528
