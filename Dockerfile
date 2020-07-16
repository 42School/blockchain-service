FROM debian:buster

RUN apt-get update -y && apt-get upgrade -y

RUN apt-get install -y nodejs npm zsh

RUN npm install -g truffle solc

COPY ./truffle-config.js .

CMD truffle develop
