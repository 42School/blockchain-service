FROM debian:buster-backports

COPY . /blockchain-service

EXPOSE 8080:8080

CMD ["/blockchain-service/blockchain-service"]