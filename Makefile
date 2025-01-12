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

.PHONY: run clean build build-arm lint lint-fix docker-build dc-up dc-down deploy aws-s3-sync publish
