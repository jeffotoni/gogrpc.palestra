FROM golang:1.8-alpine
ADD . /go/src/api.http
RUN go install api.http

FROM alpine:latest
COPY --from=0 /go/bin/api.http /
CMD ["/api.http"]