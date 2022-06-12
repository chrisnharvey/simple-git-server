FROM alpine:latest

COPY bin/simple-git-server /usr/bin/simple-git-server
COPY start.sh /start.sh
COPY etc/ /etc

RUN apk add openssh git sudo && \
    mkdir -p /git/repos && \
    adduser admin --disabled-password --shell /usr/bin/sgs --home /git && \
    adduser git --disabled-password --shell /sbin/nologin --home /git/repos

EXPOSE 22

CMD ["sh", "/start.sh"]