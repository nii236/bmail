#!/bin/bash

go run cmd/drop/main.go
go run cmd/setup/main.go
./bin/sqlboiler ./bin/sqlboiler-sqlite3 --wipe --tag db