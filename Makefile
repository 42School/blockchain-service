NAME	=	blockchain-service
CONTRACTNAME	=	FtDiploma
ETHSERV =	eth-server
APICLIENT = blockchain-service

DEPS	=	Makefile

NONE = \033[0m
RED = \033[31m
GREEN = \033[32m
YELLOW = \033[33m
BLUE = \033[34m
MAGENTA = \033[35m
CYAN = \033[36m

.PHONY:	all install testing server dev go-compile full-compile go-clean full-clean docker-stop docker-clean re

all:		install testing dev

install:
			docker build -f Dockerfile.dev -t $(APICLIENT) .

testing: server
			$(shell sleep 10)
			truffle test --network localhost

server:
			docker build -f Dockerfile.server -t $(ETHSERV) .
			docker run --add-host=$(ETHSERV):172.17.0.1 --name $(ETHSERV) -ti -p 9545:9545 -d $(ETHSERV)

dev:
			docker run --name $(APICLIENT) -ti -p 8080:8080 -d $(APICLIENT)

go-compile:
			@echo "$(YELLOW)Compiling $(NAME) in golang!$(NONE)"
			go build -o $(NAME)
			@echo "$(GREEN)$(NAME) ready!$(NONE)"

full-compile:
			@echo "$(YELLOW)Compiling the smart-contract in solidity!$(NONE)"
			truffle compile
			@echo "$(YELLOW)Compiling the smart-contract in golang!$(NONE)"
			solcjs --abi contracts/$(CONTRACTNAME).sol > $(CONTRACTNAME).abi
			solcjs --bin contracts/$(CONTRACTNAME).sol > $(CONTRACTNAME).bin
			abigen --bin=$(CONTRACTNAME).bin --abi=contracts_$(CONTRACTNAME)_sol_$(CONTRACTNAME).abi --pkg=diploma --out=$(CONTRACTNAME).go
			sed -i 's/diploma/contracts/' $(CONTRACTNAME).go
			mv $(CONTRACTNAME).go ./src/contracts/$(CONTRACTNAME).go
			@echo "$(YELLOW)Compiling $(NAME) in golang!$(NONE)"
			go build -o $(NAME)
			@echo "$(GREEN)$(NAME) ready!$(NONE)"

go-clean:
			@echo "$(YELLOW)Cleaning...$(NONE)"
			rm $(NAME)

full-clean: go-clean
			rm $(NAME).bin
			rm $(NAME).abi
			rm contracts_$(NAME)_sol_$(NAME).abi
			rm contracts_$(NAME)_sol_$(NAME).bin

docker-stop:
			docker stop $(APICLIENT)
			docker rm $(APICLIENT)

docker-rm:
			docker image rm $(APICLIENT)
			docker image rm $(ETHSERV)

docker-clean: docker-stop docker-rm

re:			docker-clean all
