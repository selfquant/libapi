import websocket
import json
try:
    import thread
except ImportError:
    import _thread as thread
import time

# // the market trigger demo for user
# /*
#   {
#     "market": {
#       "id": 1585514460,
#       "mrid": 55872638664,
#       "open": 6104.64,
#       "close": 5898.95,
#       "high": 6139,
#       "low": 5814,
#       "amount": 376281.43632446392051196659095445619343697,
#       "vol": 22875166,
#       "count": 634180
#     }
#     "lever_rate": 20,
#     "direction": "buy",
#     "cost_open": 5674.32,
#     "trade_count": 100,
#     "custom_trigger_name": "",
#     "custom_trigger_value": ""
#   }
# */

# // the trade trigger demo for user
# /*
#   {
#     "market": {
#       "ts": 1539831709001,
#       "price": 6742.25,
#       "direction": "buy"
#       "amount": 20,
#       "origin": [{
#         "amount": 20,
#         "ts": 1539831709001,
#         "id": 265842227259096443,
#         "price": 6742.25,
#         "direction": "buy"
#       }]
#     },
#     "lever_rate": 20,
#     "direction": "buy",
#     "cost_open": 5674.32,
#     "trade_count": 100
#     "custom_trigger_name": "",
#     "custom_trigger_value": ""
#   }
# */
def on_message(ws, message):
    print(message)

def on_error(ws, error):
    print(error)

def on_close(ws):
    print("### closed ###")

def on_open(ws):
    dictSub = {"userid": "123456", "market": "Huobi.market.BTC_CQ.trade.detail"}
    ws.send(json.dumps(dictSub))
    pass

if __name__ == "__main__":
    websocket.enableTrace(True)
    ws = websocket.WebSocketApp("wss://www.selfquant.com/api/subscribe?usage=trigger",
                              on_message = on_message,
                              on_error = on_error,
                              on_close = on_close)
    ws.on_open = on_open
    ws.run_forever()