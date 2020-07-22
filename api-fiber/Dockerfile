FROM golang:1.8-alpine
#ADD . /go/src/ecommerce.fiber
#RUN go install ecommerce.fiber
ADD api.fiber /go/bin/api.fiber

FROM alpine:latest
COPY --from=0 /go/bin/api.fiber /
#ENV PORT 8080
CMD ["/api.fiber"]