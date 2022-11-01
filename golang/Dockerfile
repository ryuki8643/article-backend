FROM golang:1.19.2-alpine3.16

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY ./cmd/main.go ${ROOT}

COPY go.mod ${ROOT}

RUN go mod tidy


CMD ["go", "run", "main.go"]