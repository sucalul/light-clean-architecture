build:
	docker-compose build

up:
	@make build
	docker-compose up

exec:
	docker-compose exec api sh

exec-db:
	docker-compose exec db bash -c "psql -U postgres"
