FROM golang:1.15 AS builder
RUN go get github.com/markbates/pkger/cmd/pkger

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN pkger && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o accordion

#FROM alpine:latest
FROM scratch
COPY --from=builder /etc/ssl /etc/ssl

COPY --from=builder /app/accordion /accordion
CMD ["/accordion"]