# Blockchain-service

## Description

Blockchain-service est un project blockchain de 42.

42 veut créer des diplômes stockés dans la blockchain, pour que si un jour 42 disparaît, les étudiants puissent toujours avoir un moyen de montrer leurs certifications aux entreprises.

La stack choisi pour ce projet est `go`, la blockchain `ethereum` et `solidity` pour le smartcontract.

La branche dev en cours est [dev/go/eth](https://github.com/42School/blockchain-service/tree/dev/go/eth).

La précédente version stable est la [v1.3.1](https://github.com/42School/blockchain-service/tree/v1.3.1), si vous trouvez des bugs sur cette version n'hésitez pas à faire une `issue`

## Sommaire

- [Installation](#installation)
- [Lancement en mode dev](#lancement-en-mode-dev)
- [Route de l'api](#route-de-lapi)
- [Makefile](#makefile)
- [Nouvelle Feature](#nouvelle-feature)
- [Configuration](#configuration)
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

## Nouvelle Feature

Voici les nouvelles features pour la v2:

- Une meilleure vérification lors de l'écriture d'un diplôme. Check du hash demandé lors de l'écriture et celui écrit emit par un événement blockchain.
- Une queue de retry qui ré-essaye l'écriture d'un diplôme 30 minutes après, si il avait échouer une première fois.
- Un système de roulement de compte Ethereum qui va envoyer les transactions pour écrire un diplôme en blockchain.
- Un sytème d'envoye de mail:
  - Si le seuil d'Eth est trop faible sur un compte
  - Si le système de sécurité s'active
- Un sytème de sécurité, si l'écriture d'un diplôme en blockchain est différent que celui demandé alors il s'active et envoie toutes les prochaines demande en queue de retry. Sa désactivation se fait manuellement via un mode de commande sur STDIN.
- Un système de commande, il lis des commandes sur STDIN lors de l'éxecution du programme.
  - Pour l'activer il faut écrire `cmd` dans STDIN
  - Commandes prise en charge:
    - `disable security system` pour désactiver le système de sécurité
    - `exit` pour quitté le mode de commande

## Configuration

### Roulement de compte ETH

Pour configurer le roulement des comptes Ethereum qui vont écrire sur la blockchain, vous devez créer un fichier csv tel quel:

```csv
#file name, password
UTC--2020-07-24T08-19-17.983576000Z--cac03bac6965e6d8ca96537a0344cc506b32c2c7, password
UTC--2020-07-24T08-24-31.985849000Z--fe5ac6a7bb66da6916becb74a4a3e00074cd2599, password
UTC--2020-07-24T08-25-31.194883000Z--aec7bdfb241e56c04acf5e1a2a49f147867b85b7, password
```

Puis ajouté dans l'env le path du dossier contenant les fichiers keystore:

```dockerfile
ENV KEYSTOREPATH="./keystore"
```

Pour les tests un fichier est fournis `accounts.csv` ainsi qu'un dossier `./keystore`.

### Compte Officiel de 42 

Pour configurer le compte qui va signer les diplômes vous devez créer un fichier `keystore` qui sera stocker un dossier à part du dossier pour le roulement, puis enregistrer son `path` dans l`env:

```dockerfile
ENV KEYSTOREPATHSIGN="./keystore-sign"
```

Pour les tests un dossier ainsi qu'un fichier keystore sont fournis

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
