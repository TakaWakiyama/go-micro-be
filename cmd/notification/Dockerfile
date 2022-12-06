FROM golang:1.19.1-alpine as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0


RUN apk update && apk add git
WORKDIR ${ROOT}
COPY . .

# RUN go build -o /go/bin/notification ./main.go
RUN go run main.go --subscription subscriptionname1

# RUN go mod download
EXPOSE 8082

CMD ["notification"]
