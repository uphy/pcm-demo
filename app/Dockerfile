FROM golang:1.16

WORKDIR /go/src/app

COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /bin/media ./media
RUN go build -o /bin/advertiser ./advertiser
