---
key: vnpy-connect-binance
title: vnpy-connect-binance
date: 2020-12-24 16:17:05 +0800
typora-root-url: /Users/chenjing/workspace/github/gnuser.github.io
---
使用vnpy连接binance
<!--more-->

## 新建run.py

```python
from vnpy.event import EventEngine
from vnpy.trader.engine import MainEngine
from vnpy.trader.ui import MainWindow, create_qapp
from vnpy.gateway.binance import BinanceGateway
from vnpy.app.cta_strategy import CtaStrategyApp
from vnpy.app.cta_backtester import CtaBacktesterApp

def main():
    """Start VN Trader"""
    qapp = create_qapp()

    event_engine = EventEngine()
    main_engine = MainEngine(event_engine)

    main_engine.add_gateway(BinanceGateway)
    main_engine.add_app(CtaStrategyApp)
    main_engine.add_app(CtaBacktesterApp)

    main_window = MainWindow(main_engine, event_engine)
    main_window.showMaximized()

    qapp.exec()

if __name__ == "__main__":
    main()
```





## 点击连接binance，填写api key和secret

没有注册的可以使用此链接 [https://www.binance.com/en/register?ref=UC0VE7Q5](https://www.binance.com/en/register?ref=UC0VE7Q5)，交易时，会得到10%的返佣比例

在`代码`栏中填写相应的交易对，即可展现对应的orderbook以及最新的行情数据

![image-20201224162433033](/../../../../../../../media/2020-12-24-vnpy-connect-binance/image-20201224162433033.png)

## 下单

参考上图的例子，我下了个`600`卖出`ETH`的限价单

