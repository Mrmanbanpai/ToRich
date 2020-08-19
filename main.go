package main

import (
	"ToRich/common"
	"ToRich/service"
	"flag"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"io/ioutil"

	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"
)

const (
	apisecret = "HAplcOBSdO7xJaFNtLLKTMRymniClse9Ybe5XnJzleOD5enywtI2HBItsqjNy96b"
	secret    = "Bhwi88kVgDOCuwn90cW57EfRKy6uNDt0aW11HA1t9iVjGNTwUJPcGSX1DBxkPCyt"
)

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll file for changes with this period.
	filePeriod = 10 * time.Second
)

var addr = flag.String("addr", "stream.binance.com:9443", "http service address")

var (
	//homeTempl = template.Must(template.New("").Parse(homeHTML))
	filename string
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func readFileIfModified(lastMod time.Time) ([]byte, time.Time, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, lastMod, err
	}
	if !fi.ModTime().After(lastMod) {
		return nil, lastMod, nil
	}
	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fi.ModTime(), err
	}
	return p, fi.ModTime(), nil
}

func reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func writer(ws *websocket.Conn, lastMod time.Time) {
	lastError := ""
	pingTicker := time.NewTicker(pingPeriod)
	fileTicker := time.NewTicker(filePeriod)
	defer func() {
		pingTicker.Stop()
		fileTicker.Stop()
		ws.Close()
	}()
	for {
		select {
		case <-fileTicker.C:
			var p []byte
			var err error

			p, lastMod, err = readFileIfModified(lastMod)

			if err != nil {
				if s := err.Error(); s != lastError {
					lastError = s
					p = []byte(lastError)
				}
			} else {
				lastError = ""
			}

			if p != nil {
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				if err := ws.WriteMessage(websocket.TextMessage, p); err != nil {
					return
				}
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			logrus.Println(err)
		}
		return
	}

	var lastMod time.Time
	if n, err := strconv.ParseInt(r.FormValue("lastMod"), 16, 64); err == nil {
		lastMod = time.Unix(0, n)
	}

	go writer(ws, lastMod)
	reader(ws)
}

type requestBody struct {
	//{
	//"method": "SUBSCRIBE",
	//"params":
	//[
	//"btcusdt@aggTrade",
	//"btcusdt@depth"
	//],
	//"id": 1
	//}
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int      `json:"id"`
}

func init() {
	common.InitLog(os.Stdout)

}


func main() {
	flag.Parse()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u, _ := url.Parse("wss://stream.binance.com:9443/stream?streams=btcusdt@trade/ethusdt@trade")
	//u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws/btcusdt"}
	logrus.Info(time.Now().Format(time.RFC3339), "连接到 ：", u.String())

	rich := service.NewRich(u.String())
	defer rich.C.Close()

	//go func() {
	//	defer close(done)
	//	for {
	//		_, message, err := c.ReadMessage()
	//		if err != nil {
	//			logrus.Println("read:", err)
	//			return
	//		}
	//		logrus.Printf("recv: %s", message)
	//	}
	//}()
	go rich.ReadAndPrint()

	go func() {
		ticker := time.NewTicker(time.Minute)
		t := time.NewTicker(time.Second * 5 )
		defer t.Stop()
		defer ticker.Stop()
		for {
			select {
			case <-t.C:
				time.Sleep(time.Second)
				temp := time.Since(time.Unix(service.T,0))
				if temp > time.Second * 6 {
					rich.Down <- true
				}

			case <-ticker.C:
				logrus.Info("ping================================\n\n\n")
				err := rich.Ping()
				if err != nil {
					logrus.Errorf("[main] ping 失败 , err : %v ", err)
					rich.C = service.GetConn(u.String())
					if err != nil {
						logrus.Fatal("dial:", err)
					}
				}
			}
		}
	}()
	//service.WriteMessagesage()

	for {
		select {

		case <-rich.Signal:
			time.Sleep(time.Second)
			rich.ReConnection(u.String())
			go rich.ReadAndPrint()
			logrus.Info("重新连接成功")

		case <-interrupt:
			logrus.Info("服务即将退出")
			rich.Close()
			logrus.Info("请求服务器断开链接")

			return
		}
	}

}

//LoopTable Go程一秒更新一次对应的桌子--一个Go程对应一个桌子
//func LoopTable(v *gtable.Table) {
//	defer func() {
//		err1 := recover()
//		if err1 != nil {
//			LogCommon.CommonLog.Error(fmt.Sprintf("LoopTable lock recover:[%v], tableId:[%d]", err1,v.GetTableID()))
//		}
//		v.Mu.Unlock()
//	}()
//	for {
//		// 一秒一次
//		<-time.After(time.Second)
//		v.Mu.Lock()
//
//		if v.IsRunning() {
//			v.OnTimer()
//		}
//		if v.RobotIsValid() {
//			v.RobotTimer()
//		}
//		//比赛期间退出机器人
//		v.MatchRobotExit()
//		v.Mu.Unlock()
//	}
//}

//const homeHTML = `<!DOCTYPE html>
//<html lang="en">
//    <head>
//        <title>WebSocket Example</title>
//    </head>
//    <body>
//        <pre id="fileData">{{.Data}}</pre>
//        <script type="text/javascript">
//            (function() {
//                var data = document.getElementById("fileData");
//                var conn = new WebSocket("ws://{{.Host}}/ws?lastMod={{.LastMod}}");
//                conn.onclose = function(evt) {
//                    data.textContent = 'Connection closed';
//                }
//                conn.onmessage = function(evt) {
//                    console.log('file updated');
//                    data.textContent = evt.data;
//                }
//            })();
//        </script>
//    </body>
//</html>
//
//
