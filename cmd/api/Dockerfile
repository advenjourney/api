# build stage
FROM golang:1.15.6 AS builder
COPY . /src
WORKDIR /src
RUN make clean && make build

# final stage
FROM alpine:3.13
RUN apk add --no-cache ca-certificates
WORKDIR app
COPY --from=builder /src/bin/api /app/api
ENTRYPOINT ["/app/api"]
