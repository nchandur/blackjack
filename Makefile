.PHONY: install uninstall

PREFIX := $(HOME)/blackjack
BIN_DIR := $(PREFIX)/bin
STRAT_DIR := $(PREFIX)/strategy
SAVE_FILE := $(PREFIX)/save.json

install:
	mkdir -p $(BIN_DIR)
	mkdir -p $(STRAT_DIR)

	touch $(SAVE_FILE)

	cp ./files/strategy/*.bin $(STRAT_DIR)

	go build -o blackjack
	mv blackjack $(BIN_DIR)/
	sudo cp ~/blackjack/bin/blackjack /usr/local/bin/

uninstall:
	rm -rf $(PREFIX)
	sudo rm /usr/local/bin/blackjack