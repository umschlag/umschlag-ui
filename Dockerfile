FROM alpine:edge

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf \
    /var/cache/apk/*

ADD bin/umschlag-ui /usr/bin/
ENTRYPOINT ["/usr/bin/umschlag-ui"]
CMD ["server"]
