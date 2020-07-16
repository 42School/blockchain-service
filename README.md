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
server: Lance un conteneur Docker
compile: Compile le smart-contract, convertie le smart-contract solidity en golang et compile la partie golang
clean: Supprime le binaire go et tous autres fichiers utiles à la compile et supprime le contenaire docker
fclean: Supprime l'image docker
re: Appelle les règles clean et all
```