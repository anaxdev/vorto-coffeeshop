FROM golang:1.15-alpine as builder
WORKDIR /go/src/github.com/anaxdev/go-microservice
RUN apk add -U make git
COPY . .
RUN go mod download
RUN make build

FROM alpine

COPY --from=builder /go/src/github.com/anaxdev/go-microservice/bin/server /bin/server
COPY --from=builder /go/src/github.com/anaxdev/go-microservice/bin/client /bin/client
COPY --from=builder /etc/passwd /etc/passwd

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8080
ENTRYPOINT ["/entrypoint.sh"]
