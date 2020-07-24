# Blockchain-service

## Description

Blockchain-service est un project blockchain de 42.

42 veut créer des diplômes stockés dans la blockchain, pour que si un jour 42 disparaît, les étudiants puissent toujours avoir un moyen de montrer leurs certifications aux entreprises.

La stack choisi pour ce projet est `go`, la blockchain `ethereum` et `solidity` pour le smartcontract.

La branche dev en cours est [dev/go/eth](https://github.com/lpieri/42-Alumni/tree/dev/go/eth).

## Sommaire

- [Installation](#installation)
- [Lancement en mode dev](#lancement-en-mode-dev)
- [Route de l'api](#route-de-lapi)
- [Makefile](#makefile)
- Nouvelle Feature
- [Création d'un fichier keystore](#création-dun-fichier-keystore)

## Installation

**Petite modification importante** - Avant de lancer `make install` veuillez modifier dans le Dockerfile `Dockerfile.dev` la variable `FTENDPOINT` avec l'ip du service d'alumnisation:

```dockerfile
ENV FTENDPOINT="http://[ip-42]" # By default "http://127.0.0.1:8080"
ENV VALIDATIONPATH="/[Path for validation]" # By default "/check-request"
ENV PATHRETRY="/[Path for retry]" # By default "/check-request"
```

L'API enverra:

- Toutes les 10 minutes des requêtes confirmant les diplômes en blockchain à l'adresse `EndPoint + ValidationPath` 
- Toutes les 30 minutes des requêtes validant l'inscription d'un diplôme ayant échouer sont écriture à cette adresse `EndPoint + RetryPath`

Pour lancé le projet il faut intaller `docker`

```sh
make install
```

## Lancement en mode Dev

```sh
make dev
```

## Route de l'API

Le port par défaut de l'API est `8080`.

L'API `FtDiploma` contains à ce jour 2 routes différentes:

- `/create-diploma` qui permet de crée un nouveau diplôme dans la blockchain. C'est une route `POST` qui prend comme donnée un json.
- `/get-diploma` pour vérifier si un diplôme existe en blockchain. C'est une route `POST` qui prend comme donnée un json.

Le json accepté par les deux routes est le même, voici comment il doit être formaté (un fichier template existe dans `/test/datas/template.json`):

Il doit contenir 30 skills en float ainsi que le level en float arrondie au centième.

```json
{
  "first_name": "Louise",
  "last_name": "Pieri",
  "birth_date": "1998-12-27T00:00:00Z",
  "alumni_date": "2020-06-25T00:00:00Z",
  "level": 15.17,
  "skills": [
    8.57,
    5.42,
    ...,
    4.16
  ] // (30 Skills)
}
```

## Makefile

Un makefile est fourni avec les règles suivantes:

```
all: Appelle les règles install, testing, compile
install: Build les Dockerfile
testing: Appel la règle server et lance la commande truffle test (utilise un serveur eth local ref: Dockerfile) pour tester le smart-contract
server: Lance un conteneur Docker d'un simulateur blockchain
dev: Lance le projet en mode dev dans un container docker (commande un peu lente)
compile: Compile le smart-contract, convertie le smart-contract solidity en golang et compile la partie golang
clean: Supprime le binaire go et tous autres fichiers utiles à la compilation
docker-stop: Stop les containers docker et les supprimes
docker-rm: Supprime les images docker des dockerfiles
docker-clean: Appelle les règles docker-stop et docker-rm
re: Appelle les règles docker-clean et all
```

## Création d'un fichier Keystore

Le fichier keystore est un fichier contenant votre compte eth chiffré.

Ce fichier est utilisé dans le code pour signer les diplômes. Il a pour but à terme d'être l'adresse officielle avec laquelle 42 signe les diplômes.

Pour créer un fichier keystore, il faut exécuter la commande `geth account new`

```sh
~ geth account new
INFO [07-16|15:51:50.836] Maximum peer count                       ETH=50 LES=0 total=50
Your new account is locked with a password. Please give a password. Do not forget this password.
Password: [password]
Repeat password: [password]

Your new key was generated

Public address of the key:   [public address]
Path of the secret key file: ~/Library/Ethereum/keystore/[Nom du fichier]

- You can share your public address with anyone. Others need it to interact with you.
- You must NEVER share the secret key with anyone! The key controls access to your funds!
- You must BACKUP your key file! Without the key, it's impossible to access account funds!
- You must REMEMBER your password! Without the password, it's impossible to decrypt the key!
```

### Exemple

L'adresse publique renvoyée par la commande **doit être impérativement écrite dans le smart-contract** sinon des erreurs auronts lieu.

Par exemple l'adresse publique renvoyée par `geth account new` est `0x7e12234E994384A757E2689aDdB2A463ccD3B47d`, elle devra être assigné à la variable `ftPubAddress` du contract, comme ici:

```js
pragma solidity >=0.5.8 <0.7.0;

contract	FtDiploma {

	string public constant name = "42 Alumni";
	string public constant symbol = "42A";
	string public constant linkOfRepo = "github.com/42School/blockchain-service";
	address public constant ftPubAddress = 0x7e12234E994384A757E2689aDdB2A463ccD3B47d;

  [...]
}
```

Un fichier keystore est fournie par défaut `UTC--2020-07-16T13-52-10.535505000Z--7e12234e994384a757e2689addb2a463ccd3b47d` dont le mot de passe est `password` il sert **uniquement** à faire des tests et n'as pas vocation à aller en production.
