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
func main() {
	WSClientInit("wss://www.selfquant.com", "/api/subscribe?userid=123456&usage=notify", func(from, data string) {
		fmt.Printf("%s, %s\n", from, data)
	})
}
