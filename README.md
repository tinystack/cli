# cli
Go cli Package

## 示例

### 基本示例

示例所在文件: ./examples/hello/main.go

#### 查看帮助
```shell
$ go run ./examples/hello/main.go -h

Usage: hello [global options]

this is hello command

Global Options:
  -n, --name       this is name flag
  -h, --help       Display the help information
  -v, --version    Print version information and quit
```

#### 查看版本信息
```shell
$ go run ./examples/hello/main.go --version

version: ver-0.0.1
```

#### 执行命令
```shell
$ go run ./examples/hello/main.go

flag name value is: hello world
```

#### 接收入参
```shell
$ go run ./examples/hello/main.go -n hello-world

flag name value is: hello-world
```

### 子命令嵌套

#### 查看帮助
```shell
$ go run ./examples/subcmd/main.go --help

Usage: root [global options]

this is root command

Global Options:
  -h, --help       Display the help information
  -v, --version    Print version information and quit

Available Commands:
  sub1    this is sub1 command
  sub2    this is sub2 command

Use "root [global options] <command> --help" for more information about a given command.
```

```shell
$ go run ./examples/subcmd/main.go sub2 --help

Usage: root [global options] sub2 [--options ...]

this is sub2 command

Global Options:
  -h, --help       Display the help information
  -v, --version    Print version information and quit

Options:
  -n, --name    set sub2 command's name
  -h, --help    Display the help information

Available Commands:
  sub2son    this is sub2son command

Use "root [global options] sub2 [--options ...] <command> --help" for more information about a given command.
```

#### 执行子命令

```shell
$ go run ./examples/subcmd/main.go sub2 sub2son

this is sub2son command
```