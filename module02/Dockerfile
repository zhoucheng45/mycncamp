FROM golang:1.18 AS builder
WORKDIR /build/
copy . .
ENV GOPROXY=https://goproxy.cn,direct \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0
ENV GOSUMDB=sum.golang.google.cn

RUN go build -o myhttpserver *.go


FROM alpine
WORKDIR /mygo/
COPY --from=builder /build/myhttpserver .
EXPOSE 8080
CMD ["./myhttpserver"]


