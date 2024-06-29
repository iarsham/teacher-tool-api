# Target to run dev environment
dev:
	docker-compose -f docker-compose-dev.yaml up --build -d

# Target to run production environment
up:
	docker-compose up --build -d

# Target to stop environment
down:
	docker-compose down

# Target to run logs in app
log-app:
	docker-compose logs -f app

# Target to run logs in db
log-db:
	docker-compose logs -f db
