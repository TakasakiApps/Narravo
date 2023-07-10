FROM golang:alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# RUN apk add build-base
# # RUN go build -o main .
RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-s -w" -o main .

FROM alpine

# ENV TENCENT_CLOUD_SECRET_ID=
# ENV TENCENT_CLOUD_SECRET_KEY=

WORKDIR /app

COPY --from=build /app/main /app/main

EXPOSE 8080

RUN adduser -S -D -H -h /app appuser
USER appuser

ENTRYPOINT ["/app/main"]
