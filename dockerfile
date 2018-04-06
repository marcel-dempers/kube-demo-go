FROM golang:1.9.2-alpine3.7 as builder

WORKDIR /go
COPY . .

RUN go build demoapp.go

FROM alpine:3.7

RUN mkdir /app
COPY --from=builder /go/demoapp /app/demoapp
RUN chmod +x /app/demoapp

WORKDIR /app
EXPOSE 80

CMD ["./demoapp"]
