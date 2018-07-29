FROM alpine:3.6

EXPOSE 3000
WORKDIR /go/src/github.com/yawlhead91/GalNFCPrototype
VOLUME  /go/src/github.com/yawlhead91/GalNFCPrototype

RUN apk add --update ca-certificates tzdata

COPY ./server ./server

CMD ./server
