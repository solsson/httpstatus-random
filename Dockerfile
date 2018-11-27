FROM golang:1.11.2-alpine

COPY . /go/src/knative-training.local/webserver

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go install -ldflags '-w -extldflags "-static"' \
  /go/src/knative-training.local/webserver

FROM scratch

COPY --from=0 /go/bin/webserver /webserver

ENV PORT 8080
EXPOSE $PORT

CMD ["/webserver"]
