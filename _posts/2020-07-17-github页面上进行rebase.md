---
key: github页面上进行rebase
title: github页面上进行rebase
date: 2020-07-17 11:07:51 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---

提交 PR 到github上，默认的会以 `develop`为 `base` 显示所有的 `commit` 和 `changed files`，如果我们的 PR 分支(bugfix)是基于一个还没有合并到 `develop`的分支(feature) ，而且`feature`还比较大，为了 `review` 方便，我们应该调整 `bugfix`的`base`: `develop` -> `feature`

<!--more-->

### 在 PR 页面选择 `edit，`并且调整`base branch`

![image-20200717112051077](/../../../../../../../media/2020-07-17-github页面上进行rebase/image-20200717112051077.png)



![image-20200717112025210](/../../../../../../../media/2020-07-17-github页面上进行rebase/image-20200717112025210.png)

