FROM golang:1.19 AS builder
WORKDIR /build
COPY ./main.go .
RUN go mod init server
RUN go build -o app main.go

FROM golang:1.19
COPY --from=builder /build/app /app
EXPOSE 8090
ENTRYPOINT ["/app"]
