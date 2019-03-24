SHELL = /bin/sh

srcdir = .

DATABASE_URI = postgres://user:pass@postgres-kalicoin:5432/kalicoin_test?sslmode=disable

all: build

build:
	docker run --rm -v "$PWD":/usr/src/kalicoin -w /usr/src/kalicoin golang:latest go build -v

test:
	docker run --rm \
		--name postgres-kalicoin \
		-e POSTGRES_PASSWORD=pass \
		-e POSTGRES_USER=user \
		-e POSTGRES_DB=kalicoin_test \
		-p 5432:5432 \
		-d postgres

	# Waiting for postgres to start up...
	sleep 5

	docker run --rm \
		--link postgres-kalicoin \
		-v "$$PWD":/usr/src/kalicoin \
		-w /usr/src/kalicoin \
		-e GO111MODULE=on \
		-e DATABASE_URI="postgres://user:pass@postgres-kalicoin:5432/kalicoin_test?sslmode=disable" \
		golang:latest \
		go test -mod vendor ./...

	docker stop postgres-kalicoin

clean:
	rm kalicoin
	docker stop postgres-kalicoin

.PHONY: clean
