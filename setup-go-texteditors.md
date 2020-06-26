# Installing Go and setting up a programming environment

## Installing Go

### Installing on Linux Systems

These instructions apply if you have a Linux distribution installed natively, in a virtual machine (VM) or in the Windows Subsystem for Linux (WSL).
It is advisable to use your distribution's package manager (e.g. `apt` for Ubuntu) to make it easier to maintain the installation.

#### Ubuntu-Based Systems

We recommend adding the `golang-backports` repository to get the latest version of Go, as indicated [here](https://github.com/golang/go/wiki/Ubuntu).

```sh
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
```

Then you can install the latest version of Go with the following command:

```sh
sudo apt install golang-go
```

Alternatively you can install Go with `sudo apt install golang`, though you might get an older version if you are not running the latest Ubuntu release.

#### Other Distributions

For other Linux distributions you can install Go with the distribution's package manager.

Alternatively you can [download Go from the website](https://golang.org/dl/) and follow the [instructions on the Go website](https://golang.org/doc/install) to install it manually.

### Installing on macOS

If you use the [Homebrew package manager](https://brew.sh/) you can simply install Go with the following command:

```sh
brew install go
```

Alternatively you can [download Go from the website](https://golang.org/dl/) and follow the [instructions on the Go website](https://golang.org/doc/install) to install it manually.

### Installing on Windows

We highly recommend either running Linux in a VM or WSL if you have a Windows system.
Alternatively you can [download Go from the website](https://golang.org/dl/) and follow the [instructions on the Go website](https://golang.org/doc/install) to install it manually.

## Adding Go to PATH (Linux/Mac)

You should add `$GOPATH/bin` to `PATH` such that binaries installed with `go get` can be used from the command line.
To achieve this you should add the following line to the end of the file `$HOME/.profile` (you can edit it e.g. by entering `nano $HOME/.profile` in the shell).

```sh
PATH=$PATH:$(go env GOPATH)/bin
```

The changes will take effect the next time you log in.
To make the changes take effect at once you should run the following command:

```sh
source $HOME/.profile
```

## Recommended Editors for Writing Go Code

### Visual Studio Code (VSCode)

For writing Go code (and other languages too) we highly recommend using the cross-platfrom editor **[Visual Studio Code](https://code.visualstudio.com/)**.
The teaching staff mainly use VSCode for Go programming.

#### Installing and Setting Up VSCode

Simply follow the instructions to install the program for your desired system.

For Go language support, use the **[Go extension](https://code.visualstudio.com/docs/languages/go)**, which gives you things such as intellisense (autocompletion, etc.) and linter (check code for errors).
You install the Go extension from the marketplace within VSCode.

Another useful VSCode extension is [Code Runner](https://marketplace.visualstudio.com/items?itemName=formulahendry.code-runner), which allows to run code using a keyboard shortcut or right-clicking a file instead of using ``go run``.
Runs code by default in a read-only editor.

#### Developing in WSL with VSCode

If you are developing with WSL on Windows you can use VSCode for interacting with the WSL environment.
The VSCode documentation has [detailed instructions](https://code.visualstudio.com/docs/remote/wsl) for this use case.

### GoLand

[GoLand](https://www.jetbrains.com/go/) is a proper IDE specially designed for the Go language.
This software is not free, but as a student you can create a [free student user account](https://www.jetbrains.com/community/education/?fromMenu), and thus use GoLand for free.

Some of GoLand's features include:

* Excellent refactoring support.
* On-the-fly error detection and suggestion for fixes.
* Navigation & Search.
* Run & Debug code without extra work.

### Other Editors

If you prefer some other editor there exists Go support for many editors, such as Atom, Emacs, and vim.
The Go wiki maintains a [comprehensive list](https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins) of several IDEs and text editors that can be used for Go programming.

Whichever editor you choose, it is highly recommended that you configure it to use the [goreturns](https://github.com/sqs/goreturns) tool.
This will reformat your code to follow the Go style, and make sure that all the necessary import statements are inserted (so you don’t need to write the import statement when you start using a new package.)
The goreturns tool is compatible with most editors, but may require some configuration.

Note that editors may also be able to run your code within the editor itself, but it may require some configuration.
However, using the go tool from the command line is often times preferred.
