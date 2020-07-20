NAME	=	blockchain-service
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

.PHONY:	all install testing server compile clean fclean re

all:		install testing dev

install:
			docker build -f Dockerfile.server -t $(ETHSERV) .
			docker build -f Dockerfile.dev -t $(APICLIENT) .

testing: server
			$(shell sleep 10)
			truffle test --network localhost

server:
			docker run --add-host=$(ETHSERV):172.17.0.1 --name $(ETHSERV) -ti -p 9545:9545 -d $(ETHSERV)

dev: server
			docker run --name $(APICLIENT) -ti -p 8080:8080 -d $(APICLIENT)

compile:
			@echo "$(YELLOW)Compiling the smart-contract in solidity!$(NONE)"
			truffle compile
			@echo "$(YELLOW)Compiling the smart-contract in golang!$(NONE)"
			solcjs --abi contracts/$(NAME).sol > $(NAME).abi
			solcjs --bin contracts/$(NAME).sol > $(NAME).bin
			abigen --bin=$(NAME).bin --abi=contracts_$(NAME)_sol_$(NAME).abi --pkg=diploma --out=$(NAME).go
			sed -i 's/diploma/contracts/' $(NAME).go
			mv $(NAME).go ./src/contracts/$(NAME).go
			@echo "$(YELLOW)Compiling $(NAME) in golang!$(NONE)"
			go build -o $(NAME)
			@echo "$(GREEN)$(NAME) ready!$(NONE)"

clean:
			@echo "$(YELLOW)Cleaning...$(NONE)"
			rm $(NAME)
			rm $(NAME).bin
			rm $(NAME).abi
			rm contracts_$(NAME)_sol_$(NAME).abi
			rm contracts_$(NAME)_sol_$(NAME).bin

docker-stop:
			docker stop $(ETHSERV)
			docker rm $(ETHSERV)
			docker stop $(APICLIENT)
			docker rm $(APICLIENT)

docker-clean: docker-stop
			docker image rm $(ETHSERV)
			docker image rm $(APICLIENT)

re:			docker-clean all