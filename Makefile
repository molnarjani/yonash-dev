APP_NAME = yonash-dev-server

run:
	gowebly run

docker-build:
	docker build -t ${APP_NAME}:latest -f build/Dockerfile .


docker-compose-run:
	docker-compose up -d

.PHONY: run docker-build docker-compose-up