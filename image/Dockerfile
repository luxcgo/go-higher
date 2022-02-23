FROM alpine:latest

LABEL maintainer="luxcgo <luxcgo@gmail.com>"

ENV GOPATH /home
ENV GOROOT /usr/local/go
ENV GOPROXY goproxy.cn
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk add --no-cache --virtual .build-deps wget git tar \
    && apk add --no-cache gdb binutils libc6-compat bash \
    && wget -O /tmp/go.tar.gz https://dl.google.com/go/go1.14.12.linux-arm64.tar.gz \
    && tar -C /usr/local -zxvf /tmp/go.tar.gz \
    && git clone https://github.com/go-delve/delve /tmp/delve \
    && cd /tmp/delve/cmd/dlv \
    && go build && go install \
    && rm -rf /tmp/* \
    && rm -rf $GOPATH/pkg/* \
    && rm -rf /var/cache/apk/* \
    && rm -rf /var/src/* \
    && apk del .build-deps

WORKDIR $GOPATH

STOPSIGNAL SIGQUIT