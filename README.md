
#  Blockchain-Service - Readme FR

<p align="center">
  <a href="https://github.com/42School/blockchain-service/releases"><img alt="GitHub release" src="https://img.shields.io/github/v/release/42School/blockchain-service" /></a>
  <a href="https://www.codefactor.io/repository/github/42school/blockchain-service"><img src="https://www.codefactor.io/repository/github/42school/blockchain-service/badge?s=10f20c28a71c60d44e26bb03b89762bca3792f6d" alt="CodeFactor" /></a>
</p>

---

Blockchain-service est un project blockchain de 42.

42 veut créer des diplômes stockés dans la blockchain, pour que si un jour 42 disparaît, les étudiants puissent toujours avoir un moyen de montrer leurs certifications aux entreprises.

La stack choisi pour ce projet est `go`, la blockchain `ethereum` et `solidity` pour le smartcontract.

Si vous trouvez des bugs sur cette version n'hésitez pas à faire une `issue`

Le smart-contract à été déployé sur Ropsten en test à cette [adresse](https://ropsten.etherscan.io/address/0x524e9d84B91889E5D5d8489c24E24A89e592A1e1) `0x524e9d84B91889E5D5d8489c24E24A89e592A1e1`

La branche dev en cours est [dev/go/eth](https://github.com/42School/blockchain-service/tree/dev/go/eth).

## Sommaire

- [Installation](#installation)
- [Route de l'api](#route-de-lapi)
- [Makefile](#makefile)

## Installation

### [Guide d'installation](INSTALL.md)

### [Configurer l'env](INSTALL.md#environment)

```sh
RUN_ENV="prod" # or "dev"
COMPOSE_PROFILE="prod" # or "dev"
make install
make run
```

## Route de l'API

Le port par défaut de l'API est `8080`.

L'API contains à ce jour 3 routes différentes:

- `/metrics` permet de récupérer des metriques pour Prometheus.

- `/create-diploma` permet de crée un nouveau diplôme dans la blockchain. ⚠️ Token requis

  - Méthode `POST`

  - Elle reçois le webhook d'alumnisation.

  - Elle répondra:

    - Toute suite sauf erreur:

      - ```json
      	{
          "Status":true,
          "Message":"The writing in blockchain has been done, it will be confirmed in 10 min.",
          "Data":
          {
            "Hash":"0xa613582498eb03ca0438c8e499594d03be70db2329c96b92d57faa3350658205",
            "Level":0,
            "Skills":[]
          }
        }
        ```
      
    - Toutes les 10 mins pour confirmer l'écriture d'un diplôme à l'adresse `EndPoint + ValidationPath`, sauf erreur:

      - ```json
        {
          "Status": true,
          "Message": "The 0xa613582498eb03ca0438c8e499594d03be70db2329c96b92d57faa3350658205 diploma is definitely inscribed on Ethereum.",
          "Data":
          {
            0xa613582498eb03ca0438c8e499594d03be70db2329c96b92d57faa3350658205
          }
        }
        ```

    - Toutes les 30 mins pour valider l'inscription d'un diplôme ayant échoué la 1er fois à cette adresse `EndPoint + RetryPath`:

      - ```json
        {
          "Status":true,
          "Message":"The writing in blockchain has been done, it will be confirmed in 10 min.",
          "Data":
          {
            "Hash":"0xa613582498eb03ca0438c8e499594d03be70db2329c96b92d57faa3350658205",
            "Level":0,
            "Skills":[]
          }
        }
        ```

- `/get-diploma`  vérifie si un diplôme existe en blockchain.

  - Méthode `POST`

  - Le json requis:

    - ```json
      {
        "first_name": "Louise",
        "last_name": "Pieri",
        "birth_date": "1998-12-27",
        "alumni_date": "2020-06-25"
      }
      ```

  - Elle répondra:

    - ```json
      {
        "Status":true,
        "Message":"",
        "Data":
        {
          "Hash":"",
          "Level":21.00,
          "Skills":[{"slug":"Security","level":16.42},{"slug":"Unix","level":13.87},{},{}]
        }
      }
      ```

- `/get-all-diploma` récupére tous les diplomes stockés dans la blockchain pour une éventuelle migration du contract. ⚠️ Token requis

  - Méthode `GET`

  - Elle répondra:

    - ```
      [
        {
          "Level": 2100,
          "Skills": [{"Slug":"Security","Level":1642},{"Slug":"Unix","Level":1387},...],
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
          "Skills": [{"Slug":"Security","Level":1642},{"Slug":"Unix","Level":1387},...],
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


