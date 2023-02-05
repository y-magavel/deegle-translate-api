FROM golang:1.18-alpine

WORKDIR /go/src/app
RUN apk update && apk add git
# RUN go install github.com/cosmtrek/air@v1.29.0

CMD ["go", "run", "main.go"]
# CMD ["air", "-c", ".air.toml"]
