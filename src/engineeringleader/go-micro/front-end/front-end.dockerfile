# build a small docker image
FROM alpine:latest

RUN mkdir /app

COPY frontApp /app

CMD [ "/app/frontApp" ]