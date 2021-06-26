FROM node:14.17.1-alpine

CMD ["/bin/sh"]

WORKDIR /app

RUN apk update && \
    npm install && \
    npm install -g npm && \
    npm install -g vue-cli

ADD ./ui/package*.json ./app

RUN npm install
