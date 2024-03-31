include .env.local

.PHONY: env up down lint push

env:
	cd bot && make env
	cd web && npm install

up:
	docker compose up -d --build

down:
	docker-compose down

test:
	cd bot && make test

lint:
	cd bot && make lint
	cd web && npm run check

push:
	docker login $(ACR_HOST) -u $(ACR_USER) -p $(ACR_PASSWORD)
	docker build -t bot ./bot
	docker build -t web ./web
	docker tag bot $(ACR_HOST)/bot:latest
	docker tag web $(ACR_HOST)/web:latest
	docker push $(ACR_HOST)/bot:latest
	docker push $(ACR_HOST)/web:latest
