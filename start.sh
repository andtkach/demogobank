#!/bin/bash
SERVER_ADDRESS=localhost \
SERVER_PORT=8180 \
DB_USER=root \
DB_PASSWORD=mysqlrootpass \
DB_ADDRESS=localhost \
DB_PORT=3306 \
DB_NAME=bankdb \
AUTH_SERVER=localhost:8181 \
go run main.go
