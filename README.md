
#  Blockchain-Service

<p align="center">
  <a href="https://github.com/42School/blockchain-service/releases"><img alt="GitHub release" src="https://img.shields.io/github/v/release/42School/blockchain-service" /></a>
  <a href="https://www.codefactor.io/repository/github/42school/blockchain-service"><img src="https://www.codefactor.io/repository/github/42school/blockchain-service/badge?s=10f20c28a71c60d44e26bb03b89762bca3792f6d" alt="CodeFactor" /></a>
</p>

---

Blockchain-service est un project blockchain de 42.

42 veut créer des diplômes stockés dans la blockchain, pour que si un jour 42 disparaît, les étudiants puissent toujours avoir un moyen de montrer leurs certifications aux entreprises.

La stack choisi pour ce projet est `go`, la blockchain `ethereum` et `solidity` pour le smartcontract.

La branche dev en cours est [dev/go/eth](https://github.com/42School/blockchain-service/tree/dev/go/eth).

Si vous trouvez des bugs sur cette version n'hésitez pas à faire une `issue`

Une nouvelle version du smart-contract à été redéployé sur Ropsten et est accessible à cette [adresse](https://ropsten.etherscan.io/address/0x29a5c09219a5c71a81d26922d708e472677f4548) `0x29a5c09219a5c71a81d26922d708e472677f4548`, l'ancienne version de test se trouve à l'adresse suivante `0x7dd6b2e41C3F07f16785c943B1eF6ad6eB2e34D1`

## Sommaire

- [Installation](#installation)
- [Route de l'api](#route-de-lapi)
- [Makefile](#makefile)
- [Nouvelle Feature](#nouvelle-feature)
- [Configuration](#configuration)

## Installation

**Petite modification importante** - Modifier le fichier `blockchain-service.env` plusieurs variables son à modifier dont la variable `FTENDPOINT` et l'url du service d'alumnisation:

```env
FT_END_POINT="http://[ip-42]".              # By default "http://127.0.0.1:8080"
VALIDATION_PATH="/[Path for validation]"    # By default "/check-request"
RETRY_PATH="/[Path for retry]".             # By default "/check-request"
TOKEN="0x..."                               # By default "token"
```

Pour lancé le projet il faut intaller `docker`

```sh
RUN_ENV="prod" or "dev"
make install
make run
```

## Route de l'API

Le port par défaut de l'API est `8080`.

L'API contains à ce jour 3 routes différentes:

- `/create-diploma` permet de crée un nouveau diplôme dans la blockchain. ⚠️ Token requis

  - Méthode `POST`
  - Elle reçois le webhook d'alumnisation.
  - Elle répondra:

    - Toutes les 10 minutes des requêtes confirmant les diplômes en blockchain à l'adresse `EndPoint + ValidationPath` 
    - Toutes les 30 minutes des requêtes validant l'inscription d'un diplôme ayant échouer sont écriture à cette adresse `EndPoint + RetryPath`

- `/get-diploma`  vérifie si un diplôme existe en blockchain.

  - Méthode `POST`

  - Le json requis:

    - ```
      {
        "first_name": "Louise",
        "last_name": "Pieri",
        "birth_date": "1998-12-27T00:00:00Z",
        "alumni_date": "2020-06-25T00:00:00Z"
      }
      ```

  - Elle répondra:

    - ```
      {
        "Level": 21.00
        "Skills": [14.09 10.96 9.95 9.92 9.02 7.31 7.06 6.54 5.77 5.62 3.92 2.56 2.54 2.5 2.19, ...]
      }
      ```

- `/get-all-diploma` récupére tous les diplomes stockés dans la blockchain pour une éventuelle migration du contract. ⚠️ Token requis

  - Méthode `GET`

  - Elle répondra:

    - ```
      [
        {
          "Level": 2100,
          "Skills": [857,542,620,942,661,416,902,902,902,360,222,550,145, ...],
          "Hash": [xxx,xx,xx,xx,xx,xxx,x,xx,xxx,xx,xxx,xx,xxx,xx,xx,xx,xxx,xx,xxx,xx, ...],
          "Signature":
          {
            "V": xx,
            "R": [xxx,xxx,xx,xxx,xxx,xxx,xxx,xxx,xx,xxx,xxx,xx,xxx,xx,xx,xx,xx,xx,xxx,xxx,xxx, ...],
            "S": [xxx,xxx,xx,xxx,xxx,xxx,x,xx,xxx,xx,xxx,xxx,xx,xxx,xx,xxx,xxx,xxx,xx,xxx,xxx, ...]
          }
        },
       {
          "Level": 2100,
          "Skills": [857,542,620,942,661,416,902,902,902,360,222,550,145, ...],
          "Hash": [xxx,xx,xx,xx,xx,xxx,x,xx,xxx,xx,xxx,xx,xxx,xx,xx,xx,xxx,xx,xxx,xx, ...],
          "Signature":
          {
            "V": xx,
            "R": [xxx,xxx,xx,xxx,xxx,xxx,xxx,xxx,xx,xxx,xxx,xx,xxx,xx,xx,xx,xx,xx,xxx,xxx,xxx, ...],
            "S": [xxx,xxx,xx,xxx,xxx,xxx,x,xx,xxx,xx,xxx,xxx,xx,xxx,xx,xxx,xxx,xxx,xx,xxx,xxx, ...]
          }
        },
        {...},
        {...}
      ]
      ```

## Makefile

Un makefile est fourni avec les règles suivantes:

```
all: Lance install & run
install: Compile le binaire, build l'image docker
run: Lance le projet
testing: Test le smart-contract
update-contract: Compile le smart-contract, convertie le smart-contract solidity en golang
clean: Supprime le binaire go
clean-contract: Supprime les fichiers compiler d'update-contract
docker-stop: Stop les containers docker et les supprimes
docker-rm: Supprime les images docker des dockerfiles
docker-remake: Lance docker-stop & docker-rm & run
re: Lance docker-stop & docker-rm & all
```

## Nouvelle Feature

Voici les nouvelles features pour la v2.3:

- Lis la `Webhook d'alumnisation`
- Récupère les données manquantes sur l'api de 42 (Birthdate, Level & Skills)
- Update des modules Golang

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
ENV KEYSTOREPATH="/blockchain-service/keystore"
```

Pour les tests un fichier est fournis `accounts.csv` ainsi qu'un dossier `./keystore`.

### Compte Officiel de 42 

Pour configurer le compte qui va signer les diplômes vous devez créer un fichier `keystore` qui sera stocker un dossier à part du dossier pour le roulement, puis enregistrer son `path` dans l`env:

```dockerfile
ENV KEYSTOREPATHSIGN="/blockchain-service/keystore-sign"
```

Pour les tests un dossier ainsi qu'un fichier keystore sont fournis

### Création d'un fichier Keystore

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

#### Exemple

L'adresse publique renvoyée par la commande **doit être impérativement écrite dans le smart-contract** sinon des erreurs auronts lieu.

Par exemple l'adresse publique renvoyée par `geth account new` est `0x7e12234E994384A757E2689aDdB2A463ccD3B47d`, elle devra être assigné à la variable `ftPubAddress` du contract, comme ici:

```js
pragma solidity >=0.8.0;

contract	FtDiploma {

	string public constant name = "42 Alumni";
	string public constant symbol = "42A";
	string public constant linkOfRepo = "github.com/42School/blockchain-service";
	address public constant ftPubAddress = 0x7e12234E994384A757E2689aDdB2A463ccD3B47d;

  [...]
}
```

Un fichier keystore est fournie par défaut `UTC--2020-07-16T13-52-10.535505000Z--7e12234e994384a757e2689addb2a463ccd3b47d` dont le mot de passe est `password` il sert **uniquement** à faire des tests et n'as pas vocation à aller en production.
