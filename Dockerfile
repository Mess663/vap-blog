
FROM node:latest as noder

COPY . /home/git/www/blog
WORKDIR /home/git/www/blog

RUN cd ./web \
    && npm i \
    && npm run build

FROM golang:latest

COPY --from=noder /home/git/www /www

WORKDIR /www/blog

RUN go build .
ARG MYSQL_PSW
EXPOSE 80
CMD ["bash", "/www/start_blog.sh"]



