APP_NAME = yonash-dev-server

run:
	air

lint:
	templ generate 2> /dev/null && golangci-lint run --show-stats

lint-fix:
	golangci-lint run --fix

docker-build:
	docker build -t ${APP_NAME}:latest -f build/Dockerfile .

dc-up:
	docker-compose -f build/docker-compose.yml up -d

dc-down:
	docker-compose -f build/docker-compose.yml down

.PHONY: run docker-build docker-compose-up
