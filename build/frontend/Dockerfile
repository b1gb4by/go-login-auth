FROM node:16.3.0-alpine

CMD ["/bin/sh"]

WORKDIR /app

RUN apk update && \
    npm install && \
    npm install -g npm && \
    npm install -g vue-cli

ADD ./ui/package.json /app

RUN npm install
