FROM golang:latest as builder
WORKDIR /app
ADD . /app
RUN go mod download
RUN go build -o main .
