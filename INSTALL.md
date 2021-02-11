# Blockchain-Service - Installation Guide

---

## Sommaire

- [Requirement](#requirement)
- [Keystore](#keystore)
- [Purchase of Ethereum](#purchase-of-ethereum)
- [Deployment on Ethereum](#deploye-on-ethereum)
- [Environment](#environment)
- [Run](#run)

## Requirement

* [Golang 1.15](https://golang.org/doc/install)
* [Ethereum 1.9.25](https://geth.ethereum.org/docs/install-and-build/installing-geth) // Install Ethereum, Geth, Abigen...
* Node Latest & Solcjs 0.8.0

### Linux

```shell
~ apt-get update -y && apt-get upgrade -y
~ apt-get install -y software-properties-common python3-gpg
~ add-apt-repository "deb http://ppa.launchpad.net/ethereum/ethereum/ubuntu xenial main"
~ apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 0x2A518C819BE37D2C2031944D1C52189C923F6CA9
~ apt update -y && apt upgrade -y
~ apt-get install -y ethereum
~ geth version
Geth
Version: 1.9.25...
Architecture: amd64
Go Version: go1.15.6
Operating System: linux
~ abigen --version
abigen version 1.9.25...
~ apt-get install -y curl ⚠️ Optional ⚠️
~ curl https://dl.google.com/go/go1.15.8.linux-amd64.tar.gz --output golang-1.15.tar.gz
~ tar -C /usr/local -xzf golang-1.15.tar.gz
~ export PATH=$PATH:/usr/local/go/bin
~ go version
go version go1.15.8 linux/amd64
~ apt-get install nodejs
~ npm install -g solc
~ solcjs --version
0.8.0+commit.c7dfd78e.Emscripten.clang
```

## Keystore

## Purchase of Ethereum

## Deployment on Ethereum

## Environment

## Run

