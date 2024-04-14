FROM golang:1.20

WORKDIR /app

RUN go mod init github.com/BrentGrammer/webservice

COPY . .

RUN go mod tidy

EXPOSE 4000

WORKDIR /app/cmd/api
# run go run . in the folder where main.go is to prevent passing in other files to include/compile
CMD ["go","run","."]