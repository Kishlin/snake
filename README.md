# Snake

[![Open Source](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://opensource.org/)

![version](https://img.shields.io/badge/version-2.1.3-blue)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

![Golang](https://img.shields.io/badge/Golang-1.23.5-purple)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kishlin/snake/v2)](https://goreportcard.com/report/github.com/Kishlin/snake)
[![GoDoc](https://godoc.org/github.com/Kishlin/snake?status.svg)](https://pkg.go.dev/github.com/Kishlin/snake)

![Raylib](https://img.shields.io/badge/Raylib-3.7.0-teal)

## Download

See [releases](https://github.com/Kishlin/snake/releases) to download the latest version.

## Changelog

### 2.1.3 - 2025-02-16
- Fixed a bug where the score would be saved twice under certain conditions

### 2.1.2 - 2025-02-16
- Minor performance improvement.

### 2.1.1 - 2025-02-16
- Minor performance improvement.

### 2.1.0 - 2025-02-16
- Leaderboard to keep track of scores locally.
- Apples are now circles instead of squares.
- Better instructions and helper texts.

### 2.0.0 - 2025-02-15
- Added a title screen with shortcuts and basic navigation.
- Added configurations for game speed and wall behavior.
- New score formula based on game speed and wall behavior.
- Force FPS to 60 to prevent speed issues. FPS display was removed.
- Ability to pause the game with a keyboard shortcut.
- Replaced game engine from SDL with Raylib.
- Hide the terminal on Windows.

### 1.0.0 - 2025-02-08
- Initial release

## Compile locally

### Requirements

- Go

Tested on Go 1.23.4 and 1.23.5.

### Download sources and setup dependencies
```bash
git clone git@github.com:Kishlin/snake.git
cd snake
go mod tidy
```

### Compile options

```bash
# For a specific OS
make build_linux
make build_windows_x86_64
make build_windows_i686

# To compile all at once
make build

# Or compile manually dependencing on your environment
```
