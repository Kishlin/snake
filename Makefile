MAIN_FILE=bin/main.go

.PHONY: build_linux build_windows_x86_64 build_windows_i686 run clean build rebuild

build_linux:
	@echo "Building for Linux"
	@CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -tags static -ldflags "-s -w" -o "out/snake" $(MAIN_FILE)

build_windows_x86_64:
	@echo "Building for Windows x86_64"
	@CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -tags static -ldflags "-s -w -H windowsgui" -o "out/snake.x86_64.exe" $(MAIN_FILE)

build_windows_i686:
	@echo "Building for Windows i686"
	@CGO_ENABLED=1 CC=i686-w64-mingw32-gcc GOOS=windows GOARCH=386 go build -tags static -ldflags "-s -w -H windowsgui" -o "out/snake.i686.exe" $(MAIN_FILE)

run:
	@echo "Running"
	@go run $(MAIN_FILE)

clean:
	@echo "Cleaning up"
	@rm -rf out/*

build: build_linux build_windows_x86_64 build_windows_i686

rebuild: clean build
