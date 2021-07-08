FROM golang:latest
MAINTAINER com.pulse.engineer

WORKDIR /app
COPY . /app

RUN go get github.com/swaggo/swag/cmd/swag && \
    go get -u github.com/cosmtrek/air && \
    swag init && \
    go install .

EXPOSE 8000
CMD ["air"]