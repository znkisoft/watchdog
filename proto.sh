#!/bin/bash

protoc \
  -I=proto \
  --go_out=schema \
  --go_opt=paths=source_relative \
  proto/*.proto

sleep 5

go run cmd/db/main.go