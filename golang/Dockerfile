FROM golang:1.19.2-alpine3.16

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY . ${ROOT}

RUN go mod tidy

EXPOSE 8080


CMD ["go", "run", "cmd/main.go"]