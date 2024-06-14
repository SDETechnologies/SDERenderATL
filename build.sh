#!/bin/bash
CC=x86_64-unknown-linux-gnu-gcc GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o build/main main.go
rm build/*
cp -r views build/views
cp -r static build/static
cp db build/db.db
zip -r build.zip build 