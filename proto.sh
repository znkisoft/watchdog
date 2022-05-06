#!/bin/bash

protoc \
  -I=proto \
  --go_out=schema \
  --go_opt=paths=source_relative \
  proto/*.proto