FROM golang:latest

WORKDIR /src/aitbuswebapp
ADD ./go .

RUN go mod init aitbuswebapp
RUN go get github.com/lib/pq
RUN go get github.com/gin-gonic/gin