FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go mod tidy
RUN go build -o app .

EXPOSE 8080

CMD ["./app"]
