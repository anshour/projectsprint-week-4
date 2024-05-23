FROM golang:1.22.3-alpine3.19 AS builder

WORKDIR /app

COPY . . 

RUN go mod download

RUN GOOS=linux go build -o main .

FROM alpine:3.19.1

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/app/main"]