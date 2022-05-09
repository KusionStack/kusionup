# kusionup

`kusionup` (pronounced Kusion Up) is an elegant Kusion version manager.

I want a Kusion version manager that:

* Has a minimum prerequisite to install, e.g., does not need a Kusion compiler to pre-exist.
* Is installed with a one-liner.
* Runs well on all operating systems (at least runs well on *uix as a start).
* Installs any version of Kusion and switches to it.
* Does not inject magic into your shell.
* Is written in Go.

`kusionup` is an attempt to fulfill the above features and is heavily inspired by [goup](https://github.com/owenthereal/goup).

## Installation

### One-liner

```
curl -s "http://TODO/cli/kusionup/scripts/install_kusionup.sh" | bash
```

### Manual

If you want to install manually, there are the steps:

* Download the latest `kusionup`
* Drop the `kusionup` executable to your `PATH` and make it executable: `mv kusionup /usr/local/bin/kusionup && chmod +x /usr/local/bin/kusionup`
* Add the Kusion bin directory to your shell startup script: `echo 'export PATH=$HOME/.kusionup/current/bin:$HOME/.kusionup/current/kclvm/bin:$PATH' >> ~/.bash_profile`

## Quick Start

```
$ kusionup install
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
INFO[0053] Unpacking /Users/yym/.kusionup/kusion-open@latest/kusion-darwin.tgz ... 
INFO[0059] Success: latest downloaded in /Users/yym/.kusionup/kusion-open@latest 
INFO[0059] Default Kusion is set to 'open@latest'

$ kusionup show
|    VERSION    | ACTIVE |
|---------------|--------|
|  open@latest  |   *    |

$ kusion version
```

## How it works

* `kusionup` switches to selected Kusion version.
* `kusionup default` switches to selected Kusion version.
* `kusionup init` initialize the kusionup environment file.
* `kusionup install` downloads specified version of Kusion to`$HOME/.kusionup/$SPECIFY_VERSION/` and symlinks it to `$HOME/.kusionup/current`.
* `kusionup uninstall` uninstalls the specified Kusion version.
* `kusionup reinstall` reinstalls the specified Kusion version.
* `kusionup ls-ver` lists all available Kusion versions from all Release Source.
* `kusionup show` shows the activated Kusion version located at `$HOME/.kusionup/current`.
* `kusionup version` shows the current kusionup version.

