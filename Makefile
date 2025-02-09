run:
	docker compose up -d

build:
	docker compose build

stop:
	docker compose down

clean:
	docker compose down -v --rmi all --remove-orphans
