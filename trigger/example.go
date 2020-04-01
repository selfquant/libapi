package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/jacky2478/mlog"
)

var (
	v_client *websocket.Conn
	closeCh  chan struct{}
)

func WSClientInit(strUrl, path string, recvFunc func(from, data string)) error {
	u := url.URL{Scheme: strings.Split(strUrl, "://")[0], Host: strings.Split(strUrl, "://")[1], Path: path}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		mlog.Errorf("WSClient Init failed, detail: %v", err.Error())
		return err
	}
	//defer c.Close()

	v_client = c

	closeCh = make(chan struct{}, 0)
	if recvFunc != nil {
		go receive(recvFunc)
	}
	go wait()
	return nil
}

func WSClientSend(msg string) error {
	if v_client == nil {
		return errors.New("WSClientSend failed with invalid websocket client")
	}

	if err := v_client.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		mlog.Errorf("WSClientSend failed with websocket.Send message, detail: %v", err.Error())
		return err
	}
	return nil
}

func WSClientClose() {
	if closeCh != nil {
		close(closeCh)
	}
}

func wait() {
	for {
		select {
		case <-closeCh:
			err := v_client.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				mlog.Errorf("wait failed with websocket.Close message, detail: %v", err.Error())
				return
			}
		}
	}
}

func receive(recvFunc func(from, data string)) {
	for v_client != nil {
		_, message, err := v_client.ReadMessage()
		if err != nil {
			mlog.Errorf("receive failed, detail: %v", err.Error())
			break
		}
		// mlog.Infof("receive message: %v", string(message))

		if recvFunc != nil {
			recvFunc(v_client.RemoteAddr().String(), string(message))
		}
	}

	if v_client == nil {
		return
	}
	v_client.Close()
}


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
func main() {
	WSClientInit("wss://www.selfquant.com", "/api/subscribe?usage=push", func(from, data string) {
		fmt.Printf("%s, %s\n", from, data)
	})

	dataMap := map[string]string{"userid": "123456", "market": "Huobi.market.BTC_CQ.trade.detail"}
	dataMapBuf, _ := json.Marshal(dataMap)
	WSClientSend(string(dataMapBuf))
}
