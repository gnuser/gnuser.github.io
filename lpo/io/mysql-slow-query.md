---
title: mysql-慢查询分析
permalink: /lpo/io/mysql
key: io-mysql
layout: article
sidebar:
  nav: lpo
aside:
  toc: true
---

mysql 慢查询分析

<!--more-->

## 使用`show full processlist`命令

```shell
mysql> show full processlist;
+----+------+-----------------+------+---------+------+--------------+-----------------------------------------------------+
| Id | User | Host            | db   | Command | Time | State        | Info                                                |
+----+------+-----------------+------+---------+------+--------------+-----------------------------------------------------+
| 27 | root | localhost       | test | Query   |    0 | init         | show full processlist                               |
| 28 | root | 127.0.0.1:42262 | test | Query   |    1 | Sending data | select * from products where productName='geektime' |
+----+------+-----------------+------+---------+------+--------------+-----------------------------------------------------+
2 rows in set (0.00 sec)
```

- db 表示数据库的名字；
- Command 表示 SQL 类型；
- Time 表示执行时间；
- State 表示状态；
- 而 Info 则包含了完整的 SQL 语句。

多执行几次`show full processlist`命令，观察输出

## 使用`explain`命令测试查询语句

```shell
# 切换到test库
mysql> use test;
# 执行explain命令
mysql> explain select * from products where productName='geektime';
+----+-------------+----------+------+---------------+------+---------+------+-------+-------------+
| id | select_type | table    | type | possible_keys | key  | key_len | ref  | rows  | Extra       |
+----+-------------+----------+------+---------------+------+---------+------+-------+-------------+
|  1 | SIMPLE      | products | ALL  | NULL          | NULL | NULL    | NULL | 10000 | Using where |
+----+-------------+----------+------+---------------+------+---------+------+-------+-------------+
1 row in set (0.00 sec)
```

- select_type 表示查询类型，而这里的 SIMPLE 表示此查询不包括 UNION 查询或者子查询；
- table 表示数据表的名字，这里是 products；
- type 表示查询类型，这里的 ALL 表示全表查询，但索引查询应该是 index 类型才对；
- possible_keys 表示可能选用的索引，这里是 NULL；
- key 表示确切会使用的索引，这里也是 NULL；
- rows 表示查询扫描的行数，这里是 10000。

rows 这里等于整个表的行数，说明是全表查询，尝试添加索引进行优化

```shell
mysql> CREATE INDEX products_index ON products (productName(64));
Query OK, 10000 rows affected (14.45 sec)
Records: 10000  Duplicates: 0  Warnings: 0
```
