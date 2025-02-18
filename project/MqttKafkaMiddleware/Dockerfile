FROM golang:1.23-alpine AS builder

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/main .

COPY config.yaml .

EXPOSE 8080

CMD ["./main"]


