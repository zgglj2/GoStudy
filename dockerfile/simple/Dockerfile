FROM golang

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build

WORKDIR /dist

RUN cp /build/netapp .

EXPOSE 8888

CMD ["/dist/app"]
