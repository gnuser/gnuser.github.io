---
key: postgresql字符类型使用
title: postgresql字符类型使用
date: 2020-07-03 12:37:36 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io

---

之前看了一篇文章，对数据库字段直接使用`text`类型进行了批评，那在`postgresql`中，到底应该怎么选择呢？

<!--more-->

PostgreSQL 中字符类型包括： `char`, `varchar`,  `text`

对应关系如下：

| **字符类型**                         | 描述                                       |
| ------------------------------------ | ------------------------------------------ |
| character varying(*n*), varchar(*n*) | n 为动态长度限制，存放的是动态大小的字符串 |
| character(*n*), char(*n*)            | n 为固定长度，使用空白填充剩余字节         |
| text, varchar                        | 不受长度限制的动态字符串                   |

如果使用了 `n` 限制`varchar`或者`char`的长度，那么插入超过长度的字符串会报错，有个例外是如果你超过长度部分是空格，那么`postgresql`会帮你把空格删除掉，也可以成功插入。

如果不指定`n`，直接使用`varchar`或者`character varying`，那么就和`text`一样，不受长度限制。

那到底限制的优势在哪儿呢？一般我们可能认为限制后，内存消耗或硬盘占用会更少，但实际上在`postgresql`中并不是这样，唯一的好处就是在数据库层面帮你检查一下字符串长度，并没有什么性能提升。

> There is no performance difference among these three types, apart from increased storage space when using the blank-padded type, and a few extra CPU cycles to check the length when storing into a length-constrained column. While `character(*`n`*)` has performance advantages in some other database systems, there is no such advantage in PostgreSQL; in fact `character(*`n`*)` is usually the slowest of the three because of its additional storage costs. In most situations `text` or `character varying` should be used instead.

需要注意的是，如果使用`character`和 `char`不指定`n`，那么等价于`character(1)`和`char(1)`,也就是固定长度为 1。

所以大多数情况下，我们还是可以直接在`postgresql`使用`text`和`varchar`。

