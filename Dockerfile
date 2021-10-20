FROM golang:1.16-alpine as builder

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY internal/ internal/
COPY module/ module/
COPY public/ public/
COPY main.go .

RUN go build -o main .

FROM alpine:latest as prod

WORKDIR /app

COPY --from=0 /app/main .
COPY --from=0 /app/public/ public/

CMD ["./main"]