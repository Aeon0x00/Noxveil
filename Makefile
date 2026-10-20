include .env
export

export PROJECT_ROOT=$(shell pwd)


env-up:
	docker compose up -d noxveilvpn-postgres

env-down:
	docker compose down 

env-cleanup:
	@echo "ВНИМАНИЕ: Опасность утери данных!"
	@echo -n "Очистить все volume файлы окружения? [y/N]: " && read ans && \
	if [ "$$ans" = "y" ]; then \
		docker compose down noxveilvpn-postgres && \
		rm -rf _out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Нет параметра 'seq'. Используйте: make migrate-create seq=имя_миграции"; \
		exit 1; \
	fi
	docker compose run --rm noxveilvpn-postgres-migrate \
	    create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"
migrate-action:
	docker compose run --rm noxveilvpn-postgres-migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@noxveilvpn-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		$(action)

migrate-up:
	@make migrate-action action=up
migrate-down:
	@make migrate-action action=down

run:
	@go run cmd/bot/main.go