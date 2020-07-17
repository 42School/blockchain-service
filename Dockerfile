FROM debian:buster

RUN apt-get update -y && apt-get upgrade -y

RUN apt-get install -y nodejs npm zsh

RUN npm install -g npm@latest

RUN npm install -g truffle ganache-cli solc

COPY ./truffle-config.js /.

EXPOSE 9545:9545

CMD ganache-cli --account "0x2a5c20b8657e52644b994b71d451b1a5f40188cec559e288475750710d1c54a7,500000000000000000000" -h 0.0.0.0 -p 9545 -k muirGlacier
