FROM golang:1.8-alpine
ADD . /go/src/api.rest
RUN go install api.rest

FROM alpine:latest
COPY --from=0 /go/bin/api.rest /
CMD ["/api.rest"]