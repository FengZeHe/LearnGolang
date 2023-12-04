FROM golang:alpine as builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

WORKDIR /app

COPY . .

ENV GOPROXY=https://goproxy.io

RUN go mod tidy

RUN go build -o main .

FROM alpine:3.10.0

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /app/config/prod.config.yaml /app/config/config.yaml

EXPOSE 8085

CMD ["/app/main","-env","prod"]
