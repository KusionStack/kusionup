## Introduction

[![GitHub release](https://img.shields.io/github/release/KusionStack/kusionup.svg)](https://github.com/KusionStack/kusionup/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/KusionStack/kusionup)](https://goreportcard.com/report/github.com/KusionStack/kusionup)
[![license](https://img.shields.io/github/license/KusionStack/kusionup.svg)](https://github.com/KusionStack/kusionup/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/KusionStack/kusionup.svg)](https://pkg.go.dev/github.com/KusionStack/kusionup)
[![Coverage Status](https://coveralls.io/repos/github/KusionStack/kusionup/badge.svg)](https://coveralls.io/github/KusionStack/kusionup)

> 💡 `kusionup` is a version management tool for [kusion](https://github.com/KusionStack/kusion). It is heavily inspired by [goup](https://github.com/owenthereal/goup).

## 📜️ Language

[English](https://github.com/KusionStack/kusionup/blob/main/README.md) | [简体中文](https://github.com/KusionStack/kusionup/blob/main/README-zh.md)

## ✨ Functional Overview

* support one-click installation `kusionup` through `Homebrew`, `go install`, etc.
* `kusionup` switches to selected kusion version.
* `kusionup default` switches to selected kusion version.
* `kusionup init` initialize the kusionup environment file.
* `kusionup install` downloads specified version of kusion to `$HOME/.kusionup/$SPECIFY_VERSION/` and symlinks it to `$HOME/.kusionup/current`.
* `kusionup uninstall` uninstalls the specified kusion version.
* `kusionup reinstall` reinstalls the specified kusion version.
* `kusionup ls-ver` lists all available kusion versions from all Release Source.
* `kusionup show` shows the activated kusion version located at `$HOME/.kusionup/current`.
* `kusionup version` shows the current kusionup version.

## 🛠️ Installation

### Binary (Cross-platform: windows, linux, mac ...)

To get the binary just download the latest release for your OS/Arch from the [release page](https://github.com/KusionStack/kusionup/releases) and put the binary somewhere convenient.

### Homebrew

The `KusionStack/tap` has macOS and GNU/Linux pre-built binaries available.

First installation:

```
brew install KusionStack/tap/kusionup && kusionup init
```

Upgrade:

```
brew upgrade KusionStack/tap/kusionup
```

### Script

The `kusionup` can be installed on Linux and macOS with a small install script:

```bash
curl -sSf https://raw.githubusercontent.com/KusionStack/kusionup/main/scripts/install.sh | bash
```

Or:

```bash
wget -qO- https://raw.githubusercontent.com/KusionStack/kusionup/main/scripts/install.sh | bash
```

Windows or otherwise interested users can download binaries directly from the GitHub Releases page.

### Build from Source

Starting with Go 1.17, you can install `kusionup` from source using go install:

```
go install github.com/KusionStack/kusionup@latest && kusionup init
```

### Docker

Docker users can use the following commands to pull the latest image of the `kusionup`:

```
docker pull kusionstack/kusionup:latest
```

## ⚡ Usage

```
$ kusionup init     # Need to run at first execution

$ kusionup ls-ver   # View all installable kusion versions
github@latest
github@v0.4.3
cdn@latest
cdn@v0.4.3

$ kusionup install cdn@latest   # Install the specified kusion version
Downloaded   0.0% (     2426 / 139988826 bytes) ...
Downloaded  11.4% ( 16003466 / 139988826 bytes) ...
Downloaded  21.0% ( 29433014 / 139988826 bytes) ...
Downloaded  32.2% ( 45077686 / 139988826 bytes) ...
Downloaded  41.9% ( 58642898 / 139988826 bytes) ...
Downloaded  51.2% ( 71647010 / 139988826 bytes) ...
Downloaded  61.6% ( 86258486 / 139988826 bytes) ...
Downloaded  71.2% ( 99667706 / 139988826 bytes) ...
Downloaded  81.5% (114078806 / 139988826 bytes) ...
Downloaded  91.5% (128134166 / 139988826 bytes) ...
Downloaded 100.0% (139988826 / 139988826 bytes)
INFO[0053] Unpacking ~/.kusionup/kusion-cdn@latest/kusion-darwin.tgz ... 
INFO[0059] Success: latest downloaded in ~/.kusionup/kusion-cdn@latest 
INFO[0059] Default Kusion is set to 'cdn@latest'

$ kusionup show     # View all installed kusion versions
|    VERSION    | ACTIVE |
|---------------|--------|
|   cdn@latest  |   *    |
```

For details, please refer to the [documentation](https://kusionstack.io/docs/user_docs/getting-started/install/kusionup)

## 🙏 Thanks

* [goup](https://github.com/owenthereal/goup) - Elegant Go installer
