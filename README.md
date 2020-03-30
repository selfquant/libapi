# libapi
the api service request or subscribe for user

* the notify api by websocket

``` go
// the structre for push
/*
	{
		// 用户ID
		UserID string `json:"userid"`

		// 当前触发信号
		Signal string `json:"signal"`

		// 上次触发信号
		LastSignal string `json:"last_signal"`

		// 由用户自定义订阅的触发规则
		TriggerRule string `json:"trigger_rule"`
	}
*/
// the push demo for user
/*
	{
		"userid": "123456",
		"signal": "4.73",
		"last_signal": "-0.85",
		"trigger_rule": "收益率上涨幅度大于等于5.00"
	}
*/
```

``` go
// connect:
wss://www.selfquant.com/api/subscribe?userid=123456&usage=notify
```
