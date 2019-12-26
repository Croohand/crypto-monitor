FROM golang:1.13.5

WORKDIR /go/src/github.com/Croohand/crypto-monitor
COPY . .

RUN go get -d -v ./...
RUN go test ./...
RUN go install -v ./main

ENTRYPOINT ["main"]
