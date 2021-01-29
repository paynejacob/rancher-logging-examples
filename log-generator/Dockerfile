FROM golang:1.15.7-alpine3.13 as build
WORKDIR /build
COPY go.mod .
COPY go.sum .
COPY log-generator.go .
RUN go build -o log-generator log-generator.go

FROM alpine:3.13
COPY --from=build /build/log-generator /usr/local/bin/log-generator
ENTRYPOINT ["log-generator"]