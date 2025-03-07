FROM golang:1.23.6

WORKDIR /go/src/app

COPY main.go go.mod go.sum ./

RUN go get -d ./
RUN go build -o ./pushsimple ./cmd/app

RUN adduser --disabled-password --gecos --quiet pyroscope
USER pyroscope

CMD ["./pushsimple"]