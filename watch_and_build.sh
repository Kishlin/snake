#!/bin/bash

# Define the build command
BUILD_CMD="CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -tags static -ldflags \"-s -w\" -o \"out/snake\" bin/main.go"
EXECUTABLE="snake"
OUT_DIR="out"

# Ensure inotifywait is installed
if ! command -v inotifywait &> /dev/null; then
    echo "Error: inotifywait (from inotify-tools) is required but not installed."
    exit 1
fi

# Ensure output directory exists
mkdir -p "$OUT_DIR"

# Initial build
echo "Building project..."
eval $BUILD_CMD && echo "Build successful. Running..." && (cd "$OUT_DIR" && ./$EXECUTABLE &)

# Watch for file changes
while inotifywait -e modify -e create -e delete -r ./bin/main.go ./pkg/*/*.go; do
    echo "Changes detected. Rebuilding..."
    pkill -f "$EXECUTABLE" # Kill the existing process
    eval $BUILD_CMD && echo "Build successful. Restarting..." && (cd "$OUT_DIR" && ./$EXECUTABLE &)
done

