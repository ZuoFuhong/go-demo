BINARY = blockchain

all: build test

build:
	@echo "====> Go build"
	@go build -o $(BINARY)

test:
#	./$(BINARY) createwallet
	./$(BINARY) listaddresses
#	./$(BINARY) createblockchain -address "16XaPzwmna3cyFxFjPrd6Htu4xK8Et7Rww"
#	./$(BINARY) getbalance -address "16XaPzwmna3cyFxFjPrd6Htu4xK8Et7Rww"
#	./$(BINARY) printchain
#	./$(BINARY) send -from "Luck" -to "Jack" -amount 2

.PHONY: build test
