# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

RUN mkdir /app

WORKDIR /app

ENV GOPROXY https://goproxy.cn
ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go build -o /docker-gin-dousheng

EXPOSE 8080

CMD [ "/docker-gin-dousheng" ]