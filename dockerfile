FROM golang:latest as builder
COPY . /go/src/github.com/yawlhead91/nfc_scanner_prototype/
WORKDIR /go/src/github.com/yawlhead91/nfc_scanner_prototype/


RUN go get github.com/canthefason/go-watcher && go install github.com/canthefason/go-watcher/cmd/watcher
RUN pwd
RUN go build -a -installsuffix cgo main.go
RUN pwd
RUN watcher -run github.com/yawlhead91/nfc_scanner_prototype/ -watch github.com/yawlhead91/nfc_scanner_prototype

# FROM alpine:latest  
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/
# COPY --from=0 /go/src/github.com/yawlhead91/nfc_scanner_prototype .
# CMD ["./main"]  