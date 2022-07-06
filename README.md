# cmd-reader
用来终端阅读文本文件

### 安装 

```
go get github.com/feel-easy/cmd-reader
go install github.com/feel-easy/cmd-reader
```

### 使用

##### 配置

```
Usage:
  cmd-reader config [flags]
  cmd-reader config [command]

Available Commands:
  add         添加书
  list        书列表
  remove      移除书

Flags:
  -h, --help   help for config

Global Flags:
      --config string   config file (default is .cmd-reader.yaml)

Use "cmd-reader config [command] --help" for more information about a command.
```

##### 添加书

```

Usage:
  cmd-reader config add [flags]

Flags:
  -h, --help            help for add
  -n, --name string     书名
  -p, --path string     存放路径
  -r, --remark string   备注
  -s, --schedule int    阅读进度 (default 1)

Global Flags:
      --config string   config file (default is .cmd-reader.yaml)
```

```
cmd-reader config add  -n '资本论' -p './资本论.txt' -s 100
```

##### 移除书

```
移除书

Usage:
  cmd-reader config remove [flags]

Flags:
  -h, --help          help for remove
  -n, --name string   书名

Global Flags:
      --config string   config file (default is .cmd-reader.yaml)
```

```
cmd-reader config remove  -n '资本论' 
```

##### 书列表

```
cmd-reader config list
```

##### 阅读

```
Usage:
  cmd-reader read [flags]

Flags:
  -a, --automatic   自动读
  -h, --help        help for read
  -n, --num int     要读的书序号 (default 1)
  -p, --pages int   每页展示的行数 (default 5)
  -s, --speed int   自动读的速度 (default 5)

Global Flags:
      --config string   config file (default is .cmd-reader.yaml)
```

```
cmd-reader read -n 1
```


