#!/bin/bash

# Set the environment variable and run the go build command
env SDL2_DIR="/usr/local/x86_64-w64-mingw32" \
    CGO_ENABLED="1" \
    CC="/usr/bin/x86_64-w64-mingw32-gcc" \
    GOOS="windows" \
    CGO_LDFLAGS="-L/usr/local/x86_64-w64-mingw32/lib -lmingw32 -lSDL2 -L/usr/local/x86_64-w64-mingw32/lib -LSDL2_ttf" \
    CGO_CFLAGS="-D_REENTRANT" \
    go build -x -o out/snake.exe bin/main.go

