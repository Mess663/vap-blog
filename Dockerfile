
FROM node:latest as noder

COPY . /home/git/www/blog
WORKDIR /home/git/www/blog

RUN cd ./web \
    && npm i \
    && npm run build

FROM golang:latest

COPY --from=noder /home/git/www/blog /blog

WORKDIR /blog

RUN go build .

EXPOSE 8080
CMD ["nohup", "./blog", "$MYSQL_PSW", "&"]



