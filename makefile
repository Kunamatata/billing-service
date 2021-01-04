build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down

# Example
# TEST_STRIPE_KEY=stripe_key ENV=dev make run-local
run-local:
	go run ./cmd/web