FROM alpine:3.18.4
RUN apk --no-cache add ca-certificates git
COPY tunnel /usr/local/bin/tunnel
COPY contrib/*.tpl contrib/
ENTRYPOINT ["tunnel"]
