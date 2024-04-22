# docker build --progress=plain -t standard2hsw/hpa-reporter:v0.0.1 .
FROM golang:1.22 as builder
ADD . /app
WORKDIR /app/cmd/hpa-reporter
RUN go build -o /reporter

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /reporter /reporter
CMD ["/reporter"]