FROM golang:1.23-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/main .

ENV DB_DRIVER=postgres
ENV SERVER_ADDRESS=0.0.0.0:8080

ENV GIN_MODE=release

EXPOSE 8080

CMD [ "./main" ]
