FROM golang:1.10-alpine as builder
WORKDIR /go/src/github.com/joerx/helloweb/
COPY server.go ./
COPY pkg ./pkg/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server .

FROM alpine:3.7
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY templates ./templates
COPY --from=builder /go/src/github.com/joerx/helloweb/server .
EXPOSE 8000
CMD ["./server"]
