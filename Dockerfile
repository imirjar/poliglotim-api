FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY bin/poliglotim-api .
# COPY config/config.yml ./config/config.yml
COPY .env .env

CMD ["./poliglotim-api"]