FROM golang:1.9.2-alpine3.7 as builder

WORKDIR /go
COPY . .

RUN go build demoapp.go

FROM alpine:3.7

RUN mkdir -p /usr/lib/src/
COPY --from=builder /go/demoapp /usr/lib/src/demoapp
RUN chmod +x /usr/lib/src/demoapp

WORKDIR /usr/lib/src
EXPOSE 80

CMD ["./demoapp"]