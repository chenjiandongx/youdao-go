# Golang 有道词典命令行版 [![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

> Golang 编写的有道词典命令行版本，支持单词或句子中英互译。[Python 版](https://github.com/chenjiandongx/youdao-wd)

一直觉得 Python 写一些命令行脚本还算方便，直到我知道了 `go build`....

### 构建

**获取源码**
```bash
$ git clone https://github.com/chenjiandongx/youdao-go.git
```

**编译**

* Windows
```bash
$ go get
$ go build -ldflags "-s -w" -o yd.exe cli.go core.go
```

* Linux/MacOS
```bash
$ go get
$ go build -ldflags "-s -w" -o yd cli.go core.go
```

**[upx](https://github.com/upx/upx) 压缩可执行文件**
```bash
$ upx yd    # windows 则为 upx yd.exe
```

### 使用
```
$ yd --help
yd <words> [options]

Query words meanings via the command line.

Example:
  words could be word or sentence.

  yd hello
  yd php is the best language in the world
  yd 你好

Usage:
  yd <words>...
  yd -h | --help
  yd -v | --version

Options:
  -h --help         show this help message and exit.
  -v  --version     displays the current version of youdao-go.
```

**查询单词**
```bash
$ yd coder
>>  coder: 编码器

    美:['kəudə]  英:['kəʊdə]

    n. 编码器；编码员
    n. (Coder)人名；(法)科代；(英、葡)科德尔

    coder[编码员 编码器 程序员]
    speech coder[语言编码器 音频编码器 语音编码器]
    incremental coder[增量编码器]
```

**查询句子**
```bash
$ yd talk is cheap, show me the code
>>  talk is cheap, show me the code: 说说很简单,给我代码
```

**中文翻译**
```
$ yd php是世界上最好的语言
>>  php是世界上最好的语言: PHP is one of the best language in the world
```

### LICENSE

MIT [©chenjiandongx](https://github.com/chenjiandongx)
