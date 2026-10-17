.PHONY: test build check up down up_build clean benchmark
MODULES := authentication-service broker-service listener-service logger-service mail-service
check:
	bash scripts/check.sh
test:
	@set -e; for module in $(MODULES); do (cd "$$module" && go test ./...); done
build:
	@set -e; for module in $(MODULES); do (cd "$$module" && go build ./...); done
up:
	docker compose up -d
down:
	docker compose down
up_build:
	docker compose up -d --build
clean:
	@set -e; for module in $(MODULES); do (cd "$$module" && go clean ./...); done
benchmark:
	bash scripts/benchmark.sh
