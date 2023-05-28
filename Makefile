.PHONY: up down

up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose up --build -d --remove-orphans