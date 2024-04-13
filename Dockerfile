FROM golang:1.20

WORKDIR /app

RUN go mod init github.com/BrentGrammer/webservice

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD ["go","run","cmd/api/main.go"]