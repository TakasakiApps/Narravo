FROM golang:1.20-alpine AS build

WORKDIR /app

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add build-base

COPY go.mod .
COPY go.sum .

# RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o main .

FROM alpine

WORKDIR /app

COPY --from=build /app/main /app/main

EXPOSE 8080

RUN adduser -S -D -H -h /app appuser
USER appuser

ENTRYPOINT ["/app/main"]
