# 42-Alumni

## Description

42 Alumni est un project blockchain pour 42.

42 veut créer des diplômes stockés dans la blockchain, pour que si un jour 42 disparaît, les étudiants puissent toujours avoir un moyen de montrer leurs certifications aux entreprises.

La stack choisi pour ce projet est `go`, la blockchain `ethereum` et `solidity` pour le smartcontract.

La branche dev en cours est [dev/go/eth](https://github.com/lpieri/42-Alumni/tree/dev/go/eth).

## Installation

Pour lancé le projet il faut intaller `nodejs` `npm` `geth` (`docker` uniquement en dev)

### Pour MACOS

```sh
brew tap ethereum/ethereum
brew install ethereum
brew install node npm
```

### Pour Linux

```sh
apt-get install -y nodejs npm
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install ethereum
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
	string public constant linkOfRepo = "github.com/lpieri/42-Alumni";
	address public constant ftPubAddress = 0x7e12234E994384A757E2689aDdB2A463ccD3B47d;
  
  [...]
}
```

Un fichier keystore est fournie par défaut `UTC--2020-07-16T13-52-10.535505000Z--7e12234e994384a757e2689addb2a463ccd3b47d` dont le mot de passe est `password` il sert **uniquement** à faire des tests et n'as pas vocation à aller en production.

## Makefile

Un makefile est fourni avec les règles suivantes:

```
all: Appelle les règles install, testing, compile
install: Install truffle solcjs, les modules golang et build le Dockerfile
testing: Appel la règle server et lance la commande truffle test (utilise un serveur eth local ref: Dockerfile) pour tester le smart-contract
server: Lance un conteneur Docker d'un simulateur blockchain
compile: Compile le smart-contract, convertie le smart-contract solidity en golang et compile la partie golang
clean: Supprime le binaire go et tous autres fichiers utiles à la compile et supprime le contenaire docker
fclean: Supprime l'image docker
re: Appelle les règles clean et all
```

## Lancement en mode Dev

Executé la commande `make` & `truffle migrate`

```sh
~ make
[...]
~ truffle migrate

Compiling your contracts...
===========================
> Everything is up to date, there is nothing to compile.



Starting migrations...
======================
> Network name:    'development'
> Network id:      5777
> Block gas limit: 6721975 (0x6691b7)


2_deploy_contracts.js
=====================

   Deploying 'FtDiploma'
   ---------------------
   > transaction hash:    0xa0361fdb669d2a937545c814e1ce60c1315827360f201634976b597c86b32604
   > Blocks: 0            Seconds: 0
   > contract address:    0x5df4E5590ae16ad1aC0B9A9f02A263A4C30e4d85
   > block number:        6
   > block timestamp:     1594972621
   > account:             0x3cC18ca8225c41c7ceF5CEd11AeFC1DC047f6D5D
   > balance:             99.96081982
   > gas used:            619233 (0x972e1)
   > gas price:           20 gwei
   > value sent:          0 ETH
   > total cost:          0.01238466 ETH

   > Saving artifacts
   -------------------------------------
   > Total cost:          0.01238466 ETH


Summary
=======
> Total deployments:   1
> Final cost:          0.01238466 ETH
```

Récupérer l'adresse `contract address` ici `0x5df4E5590ae16ad1aC0B9A9f02A263A4C30e4d85` et l'écrire dans le fichier `.env` à la variable `ADDRESSCONTRACT` (à la place des `...`)

```sh
#Env File Template
KEYPASSWD="password" # password of file keystore
KEYSTOREFILE="UTC--2020-07-16T13-52-10.535505000Z--7e12234e994384a757e2689addb2a463ccd3b47d" # name of principale keystore file (account of signature the diploma)
OFFICIALADDRESS="0x7e12234E994384A757E2689aDdB2A463ccD3B47d" # The address official of 42 store in the keystore
KEYSTOREPATH="./keystore" # path of dir keystore

# Network Variable
NETWORKLINK="http://127.0.0.1:9545" # truffle network
ADDRESSCONTRACT="..." # truffle migration address
RUNENV="Dev" # Dev or Prod

# Developpement variable
DEVADDRESS="0x8718F961628D97f8d20e23Ec0e264Fd51bfD6451" # truffle address
DEVPRIVATEKEY="2a5c20b8657e52644b994b71d451b1a5f40188cec559e288475750710d1c54a7" # truffle private key
```

Les variables de `developpement` n'ont pas à être modifier car le serveur docker est fait de tel sorte à que ce compte soit importé avec `500 Eth` par défaut.

Il ne reste plus qu'a lancé la commande `source .env.dev` et `./FtDiploma`

```sh
~ source .env.dev
~ ./FtDiploma
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
    6.2,
    9.42,
    6.61,
    4.16,
    9.02,
    9.02,
    9.02,
    3.6,
    2.22,
    5.5,
    1.45,
    4.35,
    2.67,
    11.22,
    7.3,
    20.6,
    10.3,
    8.17,
    11.6,
    28.28,
    9.20,
    15.7,
    11.12,
    21.26,
    3.28,
    4.23,
    20.3,
    4.16
  ]
}
```

