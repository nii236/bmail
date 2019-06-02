#!/bin/bash

# ./bin/migrate -database "sqlite://$LOCAL_DEV_DB_USER:$LOCAL_DEV_DB_PASS@$LOCAL_DEV_DB_HOST:$LOCAL_DEV_DB_PORT/$LOCAL_DEV_DB_DATABASE?sslmode=disable" -path ./migrations drop
# ./bin/migrate -database "sqlite://$LOCAL_DEV_DB_USER:$LOCAL_DEV_DB_PASS@$LOCAL_DEV_DB_HOST:$LOCAL_DEV_DB_PORT/$LOCAL_DEV_DB_DATABASE?sslmode=disable" -path ./migrations up
./bin/sqlboiler ./bin/sqlboiler-sqlite3 --wipe --tag db