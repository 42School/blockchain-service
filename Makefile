NAME	=	blockchain-service
CONTRACTNAME	=	FtDiploma
ETHSERV =	eth-server
APICLIENT = blockchain-service

NONE = \033[0m
RED = \033[31m
GREEN = \033[32m
YELLOW = \033[33m
BLUE = \033[34m
MAGENTA = \033[35m
CYAN = \033[36m

.PHONY:	all install testing run update-contract clean clean-contract docker-stop docker-clean docker-remake re

all:		install run

install:
			go mod tidy
			@echo "$(YELLOW)Compiling $(NAME) in golang!$(NONE)"
			GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(NAME)
			@echo "$(GREEN)$(NAME) ready!$(NONE)"
			docker build -f config/Dockerfile.app -t $(APICLIENT) .

run:
			docker-compose up

testing:
			docker build -f config/Dockerfile.server -t $(ETHSERV) .
			docker run --add-host=$(ETHSERV):172.17.0.1 --name $(ETHSERV) -ti -p 9545:9545 -d $(ETHSERV)
			$(shell sleep 10)
			truffle test --network localhost

update-contract:
			@echo "$(YELLOW)Compiling the smart-contract in solidity!$(NONE)"
			truffle compile
			@echo "$(YELLOW)Compiling the smart-contract in golang!$(NONE)"
			solcjs --abi contracts/$(CONTRACTNAME).sol > $(CONTRACTNAME).abi
			solcjs --bin contracts/$(CONTRACTNAME).sol > $(CONTRACTNAME).bin
			abigen --bin=contracts_$(CONTRACTNAME)_sol_$(CONTRACTNAME).bin --abi=contracts_$(CONTRACTNAME)_sol_$(CONTRACTNAME).abi --pkg=diploma --out=$(CONTRACTNAME).go
			sed -i 's/diploma/contracts/' $(CONTRACTNAME).go
			mv $(CONTRACTNAME).go ./src/dao/contracts/$(CONTRACTNAME).go

clean:
			@echo "$(YELLOW)Cleaning...$(NONE)"
			rm $(NAME)

clean-contract:
			rm $(CONTRACTNAME).bin
			rm $(CONTRACTNAME).abi
			rm contracts_$(CONTRACTNAME)_sol_$(CONTRACTNAME).abi
			rm contracts_$(CONTRACTNAME)_sol_$(CONTRACTNAME).bin

docker-stop:
			docker stop $(APICLIENT)
			docker rm $(APICLIENT)

docker-rm:
			docker image rm $(APICLIENT)

docker-remake: docker-stop docker-rm install

re:			docker-stop docker-rm all
