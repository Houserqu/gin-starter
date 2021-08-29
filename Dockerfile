FROM golang:1.16-alpine as builder

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY controller/ controller/
COPY internal/ internal/
COPY module/ module/
COPY web/ web/
COPY main.go .

RUN go build -o main .

FROM alpine:latest as prod

WORKDIR /app

COPY --from=0 /app/main .
COPY --from=0 /app/web/ web/

CMD ["./main"]