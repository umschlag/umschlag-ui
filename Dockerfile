FROM alpine:edge
MAINTAINER Thomas Boerger <thomas@webhippie.de>

EXPOSE 9000
VOLUME ["/var/lib/umschlag"]

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
    -h /var/lib/umschlag \
    -s /bin/bash \
    -G umschlag \
    -u 1000 \
    umschlag

COPY umschlag-ui /usr/bin/

ENV UMSCHLAG_UI_STORAGE /var/lib/umschlag

USER umschlag
ENTRYPOINT ["/usr/bin/umschlag-ui"]
CMD ["server"]
