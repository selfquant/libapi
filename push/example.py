import websocket
import json
try:
    import thread
except ImportError:
    import _thread as thread
import time

#   the push demo for user
# 	{
# 		"userid": "123456",
# 		"signal": "4.73",
# 		"last_signal": "-0.85",
# 		"trigger_rule": "收益率上涨幅度大于等于5.00"
# 	}
def on_message(ws, message):
    print(message)

def on_error(ws, error):
    print(error)

def on_close(ws):
    print("### closed ###")

def on_open(ws):
    pass

if __name__ == "__main__":
    websocket.enableTrace(True)
    ws = websocket.WebSocketApp("wss://www.selfquant.com/api/subscribe?userid=123456&usage=push",
                              on_message = on_message,
                              on_error = on_error,
                              on_close = on_close)
    ws.on_open = on_open
    ws.run_forever()