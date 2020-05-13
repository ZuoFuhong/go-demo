BINARY = blockchain

all: build test

build:
	@echo "====> Go Build"
	@go build -o $(BINARY)

test:
	./$(BINARY) printchain
	./$(BINARY) addblock -data "Send 1 BTC to Lucy"
	./$(BINARY) addblock -data "Pay 0.32923 BTC for a coffee"
	./$(BINARY) printchain

.PHONY: build test
