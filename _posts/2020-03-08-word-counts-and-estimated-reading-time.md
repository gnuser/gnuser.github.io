---
title: 文章字数显示以及阅读时间预估
date: 2020-03-08 16:31:40 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
tags: jekyll
---

优化阅读体验

<!--more-->

添加代码到_includes/article/top/custom.html

{% raw %}
```html
{% capture words %}
	{{ content | strip_html | strip_newlines | remove: " " | size }}
{% endcapture %}

{% capture time %}
	{{ words | divided_by: 350 | plus: 1 }} 
{% endcapture %}
<h6> {{ words }} words, {{ time }} mins </h6>
```
{% endraw %}

