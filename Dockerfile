FROM golang:1.19-alpine

WORKDIR /app

COPY go.sum ./
COPY go.mod ./
COPY *.go ./
RUN go mod download
# RUN go get github.com/gin-gonic/gin
RUN go build -o ./src

EXPOSE 8080

CMD ["/app/src"]