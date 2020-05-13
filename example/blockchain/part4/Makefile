BINARY = blockchain

all: build test

build:
	@echo "====> Go build"
	@go build -o $(BINARY)

test:
	./$(BINARY) createblockchain -address "Luck"
#	./$(BINARY) getbalance -address "Luck"
#	./$(BINARY) printchain
#	./$(BINARY) send -from "Luck" -to "Jack" -amount 2

.PHONY: build test
