FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./
RUN go mod download
RUN go build -o ./src

EXPOSE 8080

CMD ["/app/src"]