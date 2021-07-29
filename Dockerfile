# Author: Gurbaaz
FROM golang:1.16

WORKDIR $GOPATH/src/github.com/gurbaaz27/iitk-coin

COPY go.mod .
COPY go.sum .

RUN go mod download


COPY . .

RUN go build

EXPOSE 8080

CMD ["./iitk-coin"]
