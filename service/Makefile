SHELL := /bin/bash

# ==============================================================================
# Building containers
api:
	docker build \
		-f Dockerfile \
		-t api-amd64:1.0 \
		.


# ==============================================================================
# Running from within docker compose

up:
	docker-compose -f compose/compose.yaml -f compose/compose-config.yaml up --detach --remove-orphans

down:
	docker-compose -f compose/compose.yaml down --remove-orphans


# ==============================================================================
# Administration

migrate:
	go run app/admin/main.go --db-disable-tls=1 migrate

seed: migrate
	go run app/admin/main.go --db-disable-tls=1 seed	

# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f	