FROM golang:1.15.6 as builder
ENV GOPROXY="https://goproxy.cn,direct"
COPY . /go/src/hiahia
WORKDIR /go/src/hiahia
RUN make build

FROM busybox:glibc as runner
COPY --from=0 /go/src/hiahia/server /usr/local/bin/
