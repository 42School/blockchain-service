NAME	=	FtDiploma

DEPS	=	Makefile

NONE = \033[0m
RED = \033[31m
GREEN = \033[32m
YELLOW = \033[33m
BLUE = \033[34m
MAGENTA = \033[35m
CYAN = \033[36m

.PHONY:	all install compile clean re

all:		install compile

install:
			npm install -g truffle solc
			go mod download

compile:
			@echo "$(YELLOW)Compiling the smart-contract in solidity!$(NONE)"
			truffle compile
			@echo "$(YELLOW)Compiling the smart-contract in golang!$(NONE)"
			solcjs --abi contracts/FtDiplomaBase.sol > $(NAME).abi
			solcjs --bin contracts/FtDiplomaBase.sol > $(NAME).bin
			abigen --bin=$(NAME).bin --abi=contracts_FtDiplomaBase_sol_FtDiplomaBase.abi --pkg=diploma --out=$(NAME).go
			sed -i "" 's/diploma/contracts/' $(NAME).go
			mv $(NAME).go ./src/contracts/$(NAME).go
			@echo "$(YELLOW)Compiling $(NAME) in golang!$(NONE)"
			go build -o $(NAME)
			@echo "$(GREEN)$(NAME) ready!$(NONE)"

clean:
			@echo "$(YELLOW)Cleaning...$(NONE)"
			@rm $(NAME) $(NAME).bin $(NAME).abi contracts_FtDiplomaBase_sol_FtDiplomaBase.abi contracts_FtDiplomaBase_sol_FtDiplomaBase.bin

re:			clean all