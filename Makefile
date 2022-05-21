up:
	docker-compose up -d

down:
	docker-compose down

restart:
	docker-compose restart

restart-app:
	docker-compose restart app

prune:
	docker-compose down -v

logs:
	docker-compose logs --tail=120 -f

logs-app:
	docker-compose logs --tail=120 -f app

