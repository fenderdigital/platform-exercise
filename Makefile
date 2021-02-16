.DEFAULT:
	@echo 'App targets:'
	@echo '    local-image       build local image for development'
	@echo '    deps              install dependencies need to run the application through docker-compose'
	@echo '    local             spin up local environment using docker-compose'
	@echo '    migrate           migrate the local database'
	@echo '    test              run unit tests'

default: .DEFAULT

DCR = docker-compose run --rm
local-image:
	docker-compose build

deps: local-image
	$(DCR) app sh -c "go mod download && go mod vendor"

local:
	docker-compose up

migrate:
	docker-compose run --rm app go run main.go migrate up

test: setup-db
	docker-compose run --rm app go test ./repos -v \
	docker-compose down

setup-db:
	docker-compose up -d db
