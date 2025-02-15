#!/bin/bash

CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o out/snake.x86_64.exe bin/main.go
