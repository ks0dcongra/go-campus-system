FROM golang:1.20.4-alpine3.17 as builder

LABEL maintainer="Shehomebow"

RUN apk update && apk add --no-cache git
# -p選項告訴 mkdir 創建任何不存在的中間目錄
RUN mkdir -p /app

# Docker中主要工作目錄
WORKDIR /app

COPY ./go.mod ./go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# 與 ADD . /app 將目前的資料夾"."以下的所有檔案與資料夾放到container裡的/app 87分像
COPY . .

RUN go build -o andy_trainings

ENTRYPOINT  ["/app/andy_trainings"]

EXPOSE 9528
