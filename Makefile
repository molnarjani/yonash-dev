APP_NAME = yonash-dev-server
EC2_HOST = ec2-user@44.204.4.128

# Colors
CYAN = \033[0;36m
GREEN = \033[1;32m

run:
	/tmp/${APP_NAME}

run-dev:
	air

clean:
	rm -rf /tmp/${APP_NAME}

build-npm: 
	@echo "${CYAN}building static assets..."
	@npm run build

build-templ:
	@echo "${CYAN}generating go templates..."
	@templ generate

build: build-npm build-templ clean
	@echo "${CYAN}Building executable ..."
	@go build -o /tmp/${APP_NAME} cmd/server/main.go

build-amd: build-npm build-templ clean
	@echo "${CYAN}Building AMD executable ..."
	@GOARCH=amd64 GOOS=linux go build -o /tmp/${APP_NAME} cmd/server/main.go

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

# ===================
# AWS stuff
deploy: build-amd aws-s3-sync invalidate-static-distribution
	@ssh ${EC2_HOST} "rm /app/${APP_NAME}"

	@echo "${CYAN}Copying executable..."
	@scp /tmp/${APP_NAME} ${EC2_HOST}:/app/${APP_NAME}

	@echo "${CYAN}Killing old app..."
	@ssh ${EC2_HOST}  "lsof -i :8080 && lsof -i :8080 | tail -n 1 | awk '{ print $$2 }' | xargs kill || true"
	@
	@echo "${CYAN}Restarting app..."
	@ssh ${EC2_HOST} "CDN_URL='https://dos171oaztifi.cloudfront.net' nohup /app/${APP_NAME} > /tmp/${APP_NAME}.log 2>&1 &"

	@echo
	@echo "${GREEN}Deploy complete."

ssh: 
	@ssh ${EC2_HOST}

aws-s3-sync:
	@echo "Synchronizing static files to S3..."
	@aws --profile yonashdev s3 sync internal/web/static s3://yonashdev/static/

invalidate-static-distribution:
	@echo "Invalidating CloudFront distribution..."
	@aws --profile yonashdev cloudfront create-invalidation --distribution-id E2DGIBNIQVV3MB --paths "/*"


publish: docker-build aws-s3-sync
	aws --profile yonashdev ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 024848452958.dkr.ecr.us-east-1.amazonaws.com
	docker tag  yonash-dev-server:latest 024848452958.dkr.ecr.us-east-1.amazonaws.com/yonashdev:latest
	docker push 024848452958.dkr.ecr.us-east-1.amazonaws.com/yonashdev:latest

renew-cert:
	@echo "${CYAN}Renewing SSL certificate..."
	@./scripts/renew_certificate.sh

help:
	@echo "${CYAN}Available commands:${CYAN}"
	@echo "${GREEN}  run                 ${CYAN}- Run the application from /tmp"
	@echo "${GREEN}  run-dev             ${CYAN}- Run the application in development mode using air"
	@echo "${GREEN}  clean               ${CYAN}- Remove the built executable from /tmp"
	@echo "${GREEN}  build               ${CYAN}- Build the application and static assets"
	@echo "${GREEN}  build-amd           ${CYAN}- Build the application for AMD architecture"
	@echo "${GREEN}  build-npm           ${CYAN}- Build static assets using npm"
	@echo "${GREEN}  build-templ         ${CYAN}- Generate Go templates"
	@echo "${GREEN}  lint                ${CYAN}- Run linting on the codebase"
	@echo "${GREEN}  lint-fix            ${CYAN}- Fix linting issues automatically"
	@echo "${GREEN}  docker-build        ${CYAN}- Build a Docker image for the application"
	@echo "${GREEN}  dc-up               ${CYAN}- Start Docker Compose services"
	@echo "${GREEN}  dc-down             ${CYAN}- Stop Docker Compose services"
	@echo "${GREEN}  deploy              ${CYAN}- Deploy the application to AWS EC2"
	@echo "${GREEN}  ssh                 ${CYAN}- SSH into the EC2 instance"
	@echo "${GREEN}  aws-s3-sync         ${CYAN}- Sync static files to AWS S3"
	@echo "${GREEN}  invalidate-static-distribution ${CYAN}- Invalidate CloudFront distribution cache"
	@echo "${GREEN}  publish             ${CYAN}- Build and push Docker image to AWS ECR"
	@echo "${GREEN}  renew-cert          ${CYAN}- Renew the SSL certificate"

.PHONY: run clean build build-arm lint lint-fix docker-build dc-up dc-down deploy aws-s3-sync publish
