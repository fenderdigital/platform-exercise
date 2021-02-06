FROM node:12-alpine

RUN mkdir -p /service
COPY app service/app

COPY package.json service/
WORKDIR /service

RUN npm install


EXPOSE 8080


