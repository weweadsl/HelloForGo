FROM golang:alpine

ADD app bin/app

#建立log 資料夾
RUN mkdir -p /var/log/spock/

EXPOSE 8080

ENTRYPOINT ./app
