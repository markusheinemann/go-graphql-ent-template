.PHONY: help

help: ## show help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

setup_db: ## setup database for development
	./scripts/init_db.sh

migrate: export APP_ENV=dev
migrate: ## run database migrations
	go run ./cmd/migration

setup_test_db: ## setup datanase for testing
	./scripts/init_db_test.sh

setup_e2e_db: ## setup database for e2e testing
	./scripts/init_db_e2e.sh

start: ## start development server
	air