NAME	=	FtDiploma
DOCKERNAME = ftdiploma

DEPS	=	Makefile

NONE = \033[0m
RED = \033[31m
GREEN = \033[32m
YELLOW = \033[33m
BLUE = \033[34m
MAGENTA = \033[35m
CYAN = \033[36m

.PHONY:	all install testing server compile clean re

all:		install testing compile

install:
			# sudo npm install -g truffle solc
			#go mod download
			docker build -t $(DOCKERNAME) .

testing: server
			$(shell sleep 10)
			truffle test

server:
			docker run --name $(DOCKERNAME) -ti -p 9545:9545 -d $(DOCKERNAME)

compile:
			@echo "$(YELLOW)Compiling the smart-contract in solidity!$(NONE)"
			truffle compile
			@echo "$(YELLOW)Compiling the smart-contract in golang!$(NONE)"
			solcjs --abi contracts/$(NAME).sol > $(NAME).abi
			solcjs --bin contracts/$(NAME).sol > $(NAME).bin
			abigen --bin=$(NAME).bin --abi=contracts_$(NAME)_sol_$(NAME).abi --pkg=diploma --out=$(NAME).go
			sed -i "" 's/diploma/contracts/' $(NAME).go
			mv $(NAME).go ./src/contracts/$(NAME).go
			@echo "$(YELLOW)Compiling $(NAME) in golang!$(NONE)"
			go build -o $(NAME)
			@echo "$(GREEN)$(NAME) ready!$(NONE)"

clean:
			@echo "$(YELLOW)Cleaning...$(NONE)"
			@rm $(NAME) $(NAME).bin $(NAME).abi contracts_$(NAME)_sol_$(NAME).abi contracts_$(NAME)_sol_$(NAME).bin
			docker stop $(DOCKERNAME)
			docker rm $(DOCKERNAME)
            #docker image rm $(DOCKERNAME)

re:			clean all