# setup development database
setup_db:
	./scripts/init_db.sh

# setup testing database
setup_test_db:
	./scripts/init_db_test.sh

# setup testing database
setup_e2e_db:
	./scripts/init_db_e2e.sh

# start development server
start:
	air


.PHONY: setup_db start setup_test_db setup_e2e_db