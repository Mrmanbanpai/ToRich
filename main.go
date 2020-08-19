package main

import (
	"ToRich/service"
	"flag"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
	"os/signal"
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

//var (
//	//homeTempl = template.Must(template.New("").Parse(homeHTML))
//	filename string
//	upgrader = websocket.Upgrader{
//		ReadBufferSize:  1024,
//		WriteBufferSize: 1024,
//	}
//)
//
//func readFileIfModified(lastMod time.Time) ([]byte, time.Time, error) {
//	fi, err := os.Stat(filename)
//	if err != nil {
//		return nil, lastMod, err
//	}
//	if !fi.ModTime().After(lastMod) {
//		return nil, lastMod, nil
//	}
//	p, err := ioutil.ReadFile(filename)
//	if err != nil {
//		return nil, fi.ModTime(), err
//	}
//	return p, fi.ModTime(), nil
//}
//
//func reader(ws *websocket.Conn) {
//	defer ws.Close()
//	ws.SetReadLimit(512)
//	ws.SetReadDeadline(time.Now().Add(pongWait))
//	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
//	for {
//		_, _, err := ws.ReadMessage()
//		if err != nil {
//			break
//		}
//	}
//}
//
//func writer(ws *websocket.Conn, lastMod time.Time) {
//	lastError := ""
//	pingTicker := time.NewTicker(pingPeriod)
//	fileTicker := time.NewTicker(filePeriod)
//	defer func() {
//		pingTicker.Stop()
//		fileTicker.Stop()
//		ws.Close()
//	}()
//	for {
//		select {
//		case <-fileTicker.C:
//			var p []byte
//			var err error
//
//			p, lastMod, err = readFileIfModified(lastMod)
//
//			if err != nil {
//				if s := err.Error(); s != lastError {
//					lastError = s
//					p = []byte(lastError)
//				}
//			} else {
//				lastError = ""
//			}
//
//			if p != nil {
//				ws.SetWriteDeadline(time.Now().Add(writeWait))
//				if err := ws.WriteMessage(websocket.TextMessage, p); err != nil {
//					return
//				}
//			}
//		case <-pingTicker.C:
//			ws.SetWriteDeadline(time.Now().Add(writeWait))
//			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
//				return
//			}
//		}
//	}
//}
//
//func serveWs(w http.ResponseWriter, r *http.Request) {
//	ws, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		if _, ok := err.(websocket.HandshakeError); !ok {
//			logrus.Println(err)
//		}
//		return
//	}
//
//	var lastMod time.Time
//	if n, err := strconv.ParseInt(r.FormValue("lastMod"), 16, 64); err == nil {
//		lastMod = time.Unix(0, n)
//	}
//
//	go writer(ws, lastMod)
//	reader(ws)
//}
//
//type requestBody struct {
//	//{
//	//"method": "SUBSCRIBE",
//	//"params":
//	//[
//	//"btcusdt@aggTrade",
//	//"btcusdt@depth"
//	//],
//	//"id": 1
//	//}
//	Method string   `json:"method"`
//	Params []string `json:"params"`
//	ID     int      `json:"id"`
//}
//
//func init() {
//	common.InitLog(os.Stdout)
//
//}

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

	go rich.Ticker(u.String())
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
