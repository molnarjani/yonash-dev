APP_NAME = yonash-dev-server
EC2_HOST = ec2-user@54.198.152.248


run:
	air

clean:
	rm -rf /tmp/${APP_NAME}

build: clean
	go build -o /tmp/${APP_NAME} cmd/server/main.go

build-arm: clean
	GOARCH=amd64 GOOS=linux go build -o /tmp/${APP_NAME} cmd/server/main.go

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
deploy: build-arm
	ssh ${EC2_HOST} "rm /app/${APP_NAME}"

	scp /tmp/${APP_NAME} ${EC2_HOST}:/app/${APP_NAME}
	ssh ${EC2_HOST}  "lsof -i :8080 && lsof -i :8080 | tail -n 1 | awk '{ print $$2 }' | xargs kill || true"
	ssh ${EC2_HOST} "CDN_URL='https://d1jqnkgd8d6cn5.cloudfront.net' nohup /app/${APP_NAME} > /tmp/${APP_NAME}.log 2>&1 &"


aws-s3-sync:
	aws --profile eskuvoinfo s3 sync internal/web/static s3://yonash-dev/.

publish: docker-build aws-s3-sync
	aws --profile eskuvoinfo ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 024848452958.dkr.ecr.us-east-1.amazonaws.com
	docker tag  yonash-dev-server:latest 024848452958.dkr.ecr.us-east-1.amazonaws.com/yonashdev:latest
	docker push 024848452958.dkr.ecr.us-east-1.amazonaws.com/yonashdev:latest

.PHONY: run clean build build-arm lint lint-fix docker-build dc-up dc-down deploy aws-s3-sync publish