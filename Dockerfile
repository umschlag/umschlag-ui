FROM webhippie/alpine:latest

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" \
  org.label-schema.name="Umschlag UI" \
  org.label-schema.vendor="Thomas Boerger" \
  org.label-schema.schema-version="1.0"

EXPOSE 9000 80 443
VOLUME ["/var/lib/umschlag"]

ENV UMSCHLAG_UI_ASSETS /usr/share/umschlag
ENV UMSCHLAG_UI_STORAGE /var/lib/umschlag

ENTRYPOINT ["/usr/bin/umschlag-ui"]
CMD ["server"]

RUN apk add --no-cache ca-certificates mailcap

COPY dist/static /usr/share/umschlag
COPY dist/binaries/umschlag-ui-*-linux-amd64 /usr/bin/umschlag-ui
