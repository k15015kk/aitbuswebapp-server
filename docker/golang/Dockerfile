FROM golang:latest

WORKDIR /src/aitbuswebapp
ADD ./go/aitbuswebapp-api .

RUN go mod init aitbuswebapp-api
RUN go mod tidy