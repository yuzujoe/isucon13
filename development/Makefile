up-d/%:
	docker compose -f docker-compose-common.yml -f docker-compose-$*.yml up -d --build

up/%:
	docker compose -f docker-compose-common.yml -f docker-compose-$*.yml up --build

down/%:
	docker compose -f docker-compose-common.yml -f docker-compose-$*.yml down --volumes
