# syntax=docker/dockerfile:1
FROM golang:1.17.3-alpine3.14 AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /app

COPY ./ /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o secondmate .

FROM alpine:3.14

COPY --from=builder /app/secondmate /app/secondmate

RUN chmod a+x /app/secondmate

RUN apk add libcap && setcap 'cap_net_bind_service=+ep' /app/secondmate

USER nobody:nobody

CMD ["./app/secondmate"]