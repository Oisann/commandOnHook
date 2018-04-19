FROM golang:1.10-alpine as builder
WORKDIR /go/src/oisann.net/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine/git:latest
WORKDIR /root/
COPY --from=builder /go/src/oisann.net/main .
CMD ["./main"]