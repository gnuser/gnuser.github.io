---
key: mac-install-vnpy
title: mac-install-vnpy
date: 2020-12-24 15:09:05 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---
mac 安装 vnpy
<!--more-->


## 先安装pyenv

```shell
$ brew install pyenv
```

## 安装python3.7.9

```shell
$ pyenv install 3.7.9
```

## 下载项目

```shell
$ git clone https://github.com/vnpy/vnpy.git
```

## 设置环境

```shell
$ cd vnpy
$ pyenv local 3.7.9
```

## 如果python版本设置没生效

```shell
$ eval "$(pyenv init -)"
$ python -V
Python 3.7.9
```

## 直接允许安装指令

```shell
$ ./install_osx.sh
```

成功安装显示
```shell
  Building wheel for vnpy (setup.py) ... done
  Created wheel for vnpy: filename=vnpy-2.1.8-py3-none-any.whl size=84021868 sha256=9885be4620898d6a25cc505709d519421ac8460d7346d6e5cad86f39f7f03247
  Stored in directory: /private/var/folders/kf/y0zgr29s6rl0rc11r22rx39r0000gn/T/pip-ephem-wheel-cache-uh2gnbif/wheels/fe/18/91/5b6a3683234031c92698ec82866ed545ce7683d1356c45a9f4
Successfully built vnpy
Installing collected packages: vnpy
Successfully installed vnpy-2.1.8
```

## 其他安装可能遇到的问题

### brew install mongodb
```shell
➜  ~ brew install mongodb
Updating Homebrew...
==> Searching for similarly named formulae...
Error: No similarly named formulae found.
Error: No available formula or cask with the name "mongodb".
==> Searching for a previously deleted formula (in the last month)...
Error: No previously deleted formula found.
==> Searching taps on GitHub...
Error: No formulae found in taps.
```

使用以下命令安装成功
```shell
brew tap mongodb/brew
brew install mongodb-community
```

### install requirement

```shell
➜  vnpy git:(master) ✗ pip install -r requirements.txt
DEPRECATION: Python 2.7 reached the end of its life on January 1st, 2020. Please upgrade your Python as Python 2.7 is no longer maintained. pip 21.0 will drop support for Python 2.7 in January 2021. More details about Python 2 support in pip can be found at https://pip.pypa.io/en/latest/development/release-process/#python-2-support pip 21.0 will remove support for this functionality.
Defaulting to user installation because normal site-packages is not writeable
Collecting six==1.13.0
  Using cached six-1.13.0-py2.py3-none-any.whl (10 kB)
Requirement already satisfied: wheel in /Users/gnuser/Library/Python/2.7/lib/python/site-packages (from -r requirements.txt (line 2)) (0.36.2)
ERROR: Could not find a version that satisfies the requirement PyQt5==5.14.1 (from -r requirements.txt (line 3)) (from versions: none)
ERROR: No matching distribution found for PyQt5==5.14.1 (from -r requirements.txt (line 3))

```

只能pip3安装
```shell
$ pip3 install PyQt5
```

