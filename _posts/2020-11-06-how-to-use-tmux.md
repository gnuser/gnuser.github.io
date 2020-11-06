---
key: how-to-use-tmux
title: How to use tmux
date: 2020-11-06 18:41:30 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

tmux是类似 screen 的应用，可以保存运行环境，可以在退出终端后再次挂载
<!--more-->

## install

mac
```shell
brew install tmux
```

ubuntu
```shell
sudo apt-get install tmux
```

## usage

### create new session with name

```shell
tmux new -s [session-name]
```

### list all sessions

```shell
tmux ls
```

### attach

```shell
tmux attach -t [session-name]
```

### detach, can attach again

```shell
Ctrl-b d
```

### split screen vertically

```shell
Ctrl-b %
```

### split screen horizontally

```shell
Ctrl-b "
```

### switch between panes

```shell
Ctrl-b o
```



