FROM alpine:3.6

EXPOSE 3000
WORKDIR /go/src/github.com/yawlhead91/NFC_scanner_prototype
VOLUME  /go/src/github.com/yawlhead91/NFC_scanner_prototype

RUN apk add --update ca-certificates tzdata

COPY ./server ./server

CMD ./server
