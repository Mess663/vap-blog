
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
ARG MYSQL_PSW
EXPOSE 8080
CMD ["bash", "start_blog.sh"]



