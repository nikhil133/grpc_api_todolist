#!/bin/bash
$GOPATH/src/grpc_api_todolist/migrate -database "postgres://$PG_USER:$PG_PASSWORD@$PG_HOST/$PG_DB_NAME?sslmode=disable" -path "$GOPATH/src/grpc_api_todolist/db/migrations" up
echo $GOPATH/src/grpc_api_todolist/migrate -database "postgres://$PG_USER:$PG_PASSWORD@$PG_HOST/$PG_DB_NAME?sslmode=disable" -path "$GOPATH/src/grpc_api_todolist/db/migrations"