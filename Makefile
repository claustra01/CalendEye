include .env.local

.PHONY: up down push

up:
	docker compose up -d --build

down:
	docker-compose down

push:
	docker login $(ACR_HOST) -u $(ACR_USER) -p $(ACR_PASSWORD)
	docker build -t bot ./bot
	docker build -t web ./web
	docker tag bot $(ACR_HOST)/bot:latest
	docker tag web $(ACR_HOST)/web:latest
	docker push $(ACR_HOST)/bot:latest
	docker push $(ACR_HOST)/web:latest
