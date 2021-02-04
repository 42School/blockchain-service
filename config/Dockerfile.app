FROM debian:buster-backports

RUN apt-get update -y && apt-get upgrade -y

RUN apt-get install -y software-properties-common python3-gpg

RUN add-apt-repository "deb http://ppa.launchpad.net/ethereum/ethereum/ubuntu xenial main"

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 0x2A518C819BE37D2C2031944D1C52189C923F6CA9

RUN apt update -y && apt upgrade -y

RUN apt-get install -y ethereum

COPY . /blockchain-service

EXPOSE 8080:8080

CMD ["/blockchain-service/blockchain-service"]