#!/bin/bash
rm -r build/*
CC=x86_64-unknown-linux-gnu-gcc GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o build/main main.go
cp -r views build/
cp -r static build/
cp db.db build/db.db
zip -r build.zip build 
