SHELL = /bin/sh

srcdir = .

DATABASE_URI = postgres://user:pass@postgres-kalicoin:5432/kalicoin_test?sslmode=disable

all: build

build:
	docker run --rm -v "$PWD":/usr/src/kalicoin -w /usr/src/kalicoin golang:latest go build -v

test:
	docker run --rm \
		--name postgres-kalicoin \
		-e POSTGRES_USER=user \
		-e POSTGRES_DB=kalicoin_test \
		-d postgres:11

	# Waiting for postgres to start up...
	docker run --rm \
		--link postgres-kalicoin \
		-t postgres:11 \
		/bin/bash -c "while ! psql -d kalicoin_test -h postgres-kalicoin -U user -c 'select 1;'; do sleep 1; done"

	docker run --rm \
		--link postgres-kalicoin \
		-v "$$PWD":/usr/src/kalicoin \
		-w /usr/src/kalicoin \
		-e GO111MODULE=on \
		-e DATABASE_URI="postgres://user:@postgres-kalicoin:5432/kalicoin_test?sslmode=disable" \
		golang:latest \
		/bin/bash -c "go test -mod vendor ./pkg/api && go test -mod vendor ./pkg/models" 

	docker stop postgres-kalicoin

clean:
	rm kalicoin
	docker stop postgres-kalicoin

.PHONY: clean
