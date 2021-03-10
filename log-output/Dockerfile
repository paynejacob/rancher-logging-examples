FROM golang:1.15.7-alpine3.13 as build
WORKDIR /build
COPY . .
RUN go build -o log-output log-output.go

FROM alpine:3.13
COPY --from=build /build/log-output /usr/local/bin/log-output
ENTRYPOINT ["log-output"]