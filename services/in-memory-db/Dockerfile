FROM golang:1.14.1-alpine3.11

RUN apk add --no-cache git gcc libc-dev

# Install Libs needed for vscode
RUN go get golang.org/x/tools/gopls
RUN go get github.co