APP_NAME = yonash-dev-server

docker-build:
	docker build -t ${APP_NAME}:latest -f build/Dockerfile .


docker-compose-run:
	docker-compose up -d

.PHONY: docker-build docker-compose-up