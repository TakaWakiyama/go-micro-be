FROM golang:1.17.1-alpine as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0

# COPY go.mod go.sum ./
RUN apk update && apk add git
WORKDIR ${ROOT}
# RUN go mod download
EXPOSE 8080
