FROM golang AS builder
MAINTAINER com.pulse

WORKDIR /app
COPY . /app

RUN go get github.com/swaggo/swag/cmd/swag && \
    swag init && \
    go install .

EXPOSE 8000