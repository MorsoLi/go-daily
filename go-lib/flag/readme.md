## Go 标准库提供的命令行选项解析
### 使用flag库的一般步骤：

1. 定义一些全局变量存储选项的值，如这里的intflag/boolflag/stringflag；
2. 在init方法中使用flag.TypeVar方法定义选项，这里的Type可以为基本类型Int/Uint/Float64/Bool，还可以是时间间隔time.Duration。定义时传入变量的地址、选项名、默认值和帮助信息；
3. 在main方法中调用flag.Parse从os.Args[1:]中解析选项。因为os.Args[0]为可执行程序路径，会被剔除。

### 选项格式
```bash
-arg
-arg=argvalue
-arg argvalue
```
* 遇到第一个非选项参数（即不是以-和--开头的）或终止符--，解析停止
```bash
cmd-tool -- -arg1=value1
cmd-tool arg1 --arg2=value2
```

### 