FROM golang:latest
WORKDIR /home/git/www/blog
ADD . /home/git/www/blog
RUN go build .
EXPOSE 8080
ENTRYPOINT ["./blog"]