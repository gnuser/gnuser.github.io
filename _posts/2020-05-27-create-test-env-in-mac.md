---
key: create-test-env-in-mac
title: Create test env in mac
date: 2020-05-27 22:01:08 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

在 mac 上安装 ubuntu 测试环境

<!--more-->

## 安装 docker

https://download.docker.com/mac/stable/Docker.dmg

## 获取 ubuntu 镜像

```Bash
sudo docker pull ubuntu
```

## 创建 ubuntu容器

并启动bash交互终端

```bash
sudo docker run -i -t --name mineos ubuntu bash
```

## 更新软件源

```bash
apt-get update
```

## 安装 vim

```bash
apt install vim
```

## 安装 ssh

```bash
apt-get install openssh-server
```

## 配置 sshd

```bash
vim /etc/ssh/sshd_config
# 添加下列三行
PermitRootLogin yes
PubkeyAuthentication yes
AuthorizedKeysFile	.ssh/authorized_keys
```

## 重启sshd

```bash
/etc/init.d/ssh restart
```

## 添加 mac 本机的ssh 公钥

```bash
mkdir ~/.ssh
vim ~/.ssh/authorized_keys # 粘贴 mac 本机~/.ssh/id_rsa.pub 文件内容
```

## 更新镜像

获取 mineos 的container id

```bash
sudo docker ps -a
➜ docker ps -a
CONTAINER ID        IMAGE                                   COMMAND                  CREATED             STATUS                          PORTS                                                                              NAMES
04f9076bc4ae        ubuntu                                  "bash"                   8 minutes ago       Exited (0) 11 seconds ago
```

提交

```bash
docker commit -m 'add ssh' -a 'gnuser' 04f9076bc4ae ubuntu-ssh
```

## 删除mineos 镜像

```bash
sudo docker rm mineos
```

## 启动新镜像，开启 ssh 服务

```bash
sudo docker run -d -p 26122:22 --name learn ubuntu-ssh /usr/sbin/sshd -D
```

## mac 使用ssh登陆

```bash
ssh -p 26122 root@localhost
```

可以多次登陆，进行性能测试了。