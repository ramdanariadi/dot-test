FROM golang:1.18.9-alpine as BUILD

WORKDIR $GOPATH/src/github.com/ramdanariadi/dot-test

COPY . .
RUN go mod download
RUN go build -o /app

EXPOSE 8080
ENTRYPOINT [ "/app" ]