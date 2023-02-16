FROM golang:1.18-alpine

WORKDIR /go/src/app
RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]