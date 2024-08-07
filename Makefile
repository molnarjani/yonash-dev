APP_NAME = yonash-dev-server

run:
	gowebly run

lint:
	templ generate 2> /dev/null && golangci-lint run --show-stats

lint-fix:
	golangci-lint run --fix

docker-build:
	docker build -t ${APP_NAME}:latest -f build/Dockerfile .


docker-compose-run:
	docker-compose up -d

.PHONY: run docker-build docker-compose-up
