# libapi
the api service request or subscribe for user

* the push api by websocket

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
wss://www.selfquant.com/api/subscribe?usage=push
```

* the trigger api by websocket

``` go
// the market trigger demo for user
/*
	{
		"market": {
			"id": 1585514460,
			"mrid": 55872638664,
			"open": 6104.64,
			"close": 5898.95,
			"high": 6139,
			"low": 5814,
			"amount": 376281.43632446392051196659095445619343697,
			"vol": 22875166,
			"count": 634180
		}
		"lever_rate": 20,
		"direction": "buy",
		"cost_open": 5674.32,
		"trade_count": 100,
		"custom_trigger_name": "",
		"custom_trigger_value": ""
	}
*/

// the trigger demo for user
/*
	{
		"market": {
			"ts": 1539831709001,
			"price": 6742.25,
			"direction": "buy"
			"amount": 20,
			"origin": [{
				"amount": 20,
				"ts": 1539831709001,
				"id": 265842227259096443,
				"price": 6742.25,
				"direction": "buy"
			}]
		},
		"lever_rate": 20,
		"direction": "buy",
		"cost_open": 5674.32,
		"trade_count": 100
		"custom_trigger_name": "",
		"custom_trigger_value": ""
	}
*/
```

``` go
// connect:
wss://www.selfquant.com/api/subscribe?usage=push
```