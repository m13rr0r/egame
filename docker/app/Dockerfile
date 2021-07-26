FROM golang:1.16.0-alpine3.13

RUN apk add --update \
    build-base \
    tzdata

ENV TZ=Europe/Moscow

WORKDIR /go/src/egame

COPY ./app/go.mod ./app/go.sum ./

RUN go mod download

COPY ./app .

RUN go build -o main .

CMD ["go","run","main.go"]