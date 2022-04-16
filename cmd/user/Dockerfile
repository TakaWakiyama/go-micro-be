FROM golang:1.17.1-alpine as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0


RUN apk update && apk add git
WORKDIR ${ROOT}
COPY . .

RUN go build -o /go/bin/user ./main.go
# RUN go mod download
EXPOSE 8081

CMD [ "user" ]
