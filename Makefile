.PHONY: build debug run db clean help

build:
	cd orderbook ; go build

run:
	cd orderbook ; go run .

db:
	cd tarantool ; tarantool -i main.lua

clean:
	rm -f tarantool/*.snap
	rm -f tarantool/*.xlog
	rm -f orderbook/orderbook

help:
	@printf "Makefile usage:\n"
	@printf "\n"
	@printf "\033[32m  build \033[0m compile the program\n"
	@printf "\033[32m  run   \033[0m run the program\n"
	@printf "\033[32m  db    \033[0m launch the database daemon\n"
	@printf "\033[32m  clean \033[0m remove all artifacts and logs\n"
	@printf "\033[32m  help  \033[0m display this message\n"
	@printf "\n"
