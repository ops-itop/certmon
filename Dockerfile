FROM alpine:3.8

RUN apk add --no-cache ca-certificates
ADD bin/certmon /bin/
ADD assets /assets
COPY init.sh /

CMD ["sh", "init.sh"]
