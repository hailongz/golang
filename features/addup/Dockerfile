FROM hailongz/kk-golang:latest

COPY ./main /bin/kk-addup

RUN chmod +x /bin/kk-addup

COPY ./app.json /app.josn

EXPOSE 80

ENV KK_CONFIG_FILE /app.json

ENV KK_NODE node-0

ENV KK_VERSION 1.0

WORKDIR /

CMD kk-addup 
