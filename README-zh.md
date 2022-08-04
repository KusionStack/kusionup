## 简介

[![GitHub release](https://img.shields.io/github/release/KusionStack/kusionup.svg)](https://github.com/KusionStack/kusionup/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/KusionStack/kusionup)](https://goreportcard.com/report/github.com/KusionStack/kusionup)
[![license](https://img.shields.io/github/license/KusionStack/kusionup.svg)](https://github.com/KusionStack/kusionup/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/KusionStack/kusionup.svg)](https://pkg.go.dev/github.com/KusionStack/kusionup)
[![Coverage Status](https://coveralls.io/repos/github/KusionStack/kusionup/badge.svg)](https://coveralls.io/github/KusionStack/kusionup)

> 💡 `kusionup` 是一个针对 [kusion](https://github.com/KusionStack/kusion) 的版本管理工具，它深受 [goup](https://github.com/owenthereal/goup) 的启发

## 📜️ 语言

[English](https://github.com/KusionStack/kusionup/blob/main/README.md) | [简体中文](https://github.com/KusionStack/kusionup/blob/main/README-zh.md)

## ✨ 功能简介

* 支持通过 `Homebrew`, `go install` 等一键安装 `kusionup`
* `kusionup` 切换不同的 kusion 版本
* `kusionup default` 切换指定的 kusion 版本
* `kusionup init` 初始化环境变量文件
* `kusionup install` 下载指定的 kusion 版本到 `$HOME/.kusionup/$SPECIFY_VERSION/`，然后软链接到 `$HOME/.kusionup/current`
* `kusionup uninstall` 卸载指定的 kusion 版本
* `kusionup reinstall` 重新安装指定的 kusion 版本
* `kusionup ls-ver` 列出所有可用的 kusion 版本
* `kusionup show` 展示当前安装的所有版本和当前激活版本
* `kusionup version` 展示当前 kusionup 的版本

## 🛠️ 安装

### 二进制安装（跨平台: windows, linux, mac ...）

从二进制安装，只需从 `kusionup` 的 [发布页面](https://github.com/KusionStack/kusionup/releases) 下载对应平台的二进制文件，然后将二进制文件放在命令行能访问到的目录中即可。

### Homebrew

`KusionStack/tap` 有 MacOS 和 GNU/Linux 的预编译二进制版本可用。

第一次安装:

```
brew install KusionStack/tap/kusionup && kusionup init
```

升级:

```
brew upgrade KusionStack/tap/kusionup
```

### 脚本安装

在 Linux 和 MacOS 环境中，`kusionup` 可以通过脚本一键安装：

```bash
curl -sSf https://raw.githubusercontent.com/KusionStack/kusionup/main/scripts/install.sh | bash
```

Or:

```bash
wget -qO- https://raw.githubusercontent.com/KusionStack/kusionup/main/scripts/install.sh | bash
```

Windows 或者其它感兴趣的用户可以直接在 Github Release 页面中下载二进制文件。

### 从源码构建

使用 Go 1.17+ 版本，你可以通过 `go install` 直接从源码安装 `kusionup`：

```
go install github.com/KusionStack/kusionup/cmd@latest && kusionup init
```

### Docker

Docker 用户可以用以下命令拉取 `kusionup` 的镜像：

```
docker pull kusionstack/kusionup:latest
```

## ⚡ 使用

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

详情请参考[文档](https://kusionstack.io/docs/user_docs/getting-started/install/kusionup)

## 🙏 感谢

* [goup](https://github.com/owenthereal/goup) - Elegant Go installer
