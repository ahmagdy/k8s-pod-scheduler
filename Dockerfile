FROM golang:1.14.3 AS builder

WORKDIR /app
ADD . /app
RUN go get -v -t -d ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./example/...

FROM alpine:latest AS production

COPY --from=builder /app .
LABEL maintainer="Ahmed Magdy(magdy.dev)"
CMD ["./main"]
EXPOSE 8080