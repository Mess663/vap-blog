FROM golang:latest
WORKDIR /home/git/www/blog
ADD . /home/git/www/blog

RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.3/install.sh | bash
RUN export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")" [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
RUN nvm i node
RUN nvm use node
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./blog"]
CMD ["deploy.sh"]

