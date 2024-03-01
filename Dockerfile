FROM golang:1.21.3-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o tx-relay

ENTRYPOINT [ "tx-relay" ]