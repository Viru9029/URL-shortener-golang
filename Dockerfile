# Dockerfile
FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o url-shortener

CMD ["./url-shortener"]

EXPOSE 8080
