FROM alpine:edge
MAINTAINER Thomas Boerger <thomas@webhippie.de>

EXPOSE 9000

RUN apk update && \
  apk add \
    ca-certificates \
    bash && \
  rm -rf \
    /var/cache/apk/* && \
  addgroup \
    -g 1000 \
    umschlag && \
  adduser -D \
    -h /home/umschlag \
    -s /bin/bash \
    -G umschlag \
    -u 1000 \
    umschlag

COPY bin/umschlag-ui /usr/bin/

USER umschlag
ENTRYPOINT ["/usr/bin/umschlag-ui"]
CMD ["server"]
