#!/bin/bash

CGO_ENABLED=1 CC=i686-w64-mingw32-gcc GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o out/snake.i686.exe bin/main.go
