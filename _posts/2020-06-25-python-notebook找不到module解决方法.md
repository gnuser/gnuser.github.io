---
key: python-notebook找不到module解决方法
title: python notebook找不到module解决方法
date: 2020-06-25 11:21:28 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

比如`pandas`,`matplotlib`找不到，提示`ModuleNotFoundError: No module named 'matplotlib'`

<!--more-->

1. 获取 python 路径

```shell
import sys
print(sys.executable) 
```

比如我这里是

```
/Users/gnuser/miniconda3/bin/python
```

2. 使用该路径进行包安装

```shell
/Users/gnuser/miniconda3/bin/python -m pip install pandas
/Users/gnuser/miniconda3/bin/python -m pip install matplotlib
```

