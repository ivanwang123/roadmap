FROM golang:1.16-alpine

WORKDIR /go/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get github.com/cosmtrek/air

COPY .air.toml .

EXPOSE 8080

CMD ["air"]