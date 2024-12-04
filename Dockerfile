FROM golang:1.23.3

WORKDIR /app

COPY . .

RUN go build -o honeypot main.go

EXPOSE 2222

CMD ["./honeypot"]
