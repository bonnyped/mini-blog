#!/bin/bash

export DB_STRING='user=postgres password=123 dbname=blog sslmode=disable'
export GOOSE_DRIVER='postgres'
export GOOSE_DBSTRING='postgres://postgres:123@localhost:5432/blog?sslmode=disable'
export GOOSE_MIGRATION_DIR='./migrations'
