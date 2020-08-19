package service

import (
	"ToRich/repository/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"runtime"
	"strconv"
	"time"
)
var T int64

type Rich struct {
	C      *websocket.Conn
	Signal chan bool
	Down chan bool
}

func NewRich(url string) *Rich {
	return &Rich{
		C:      GetConn(url),
		Signal: make(chan bool),
		Down: make(chan bool),
	}
}

func GetConn(url string) *websocket.Conn {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		logrus.Fatal("dial:", err)
	}
	return c
}

func (r *Rich) ReConnection(url string) {
	err := r.C.Close()
	if err != nil {
		logrus.Errorf("重连断开失败 err : %v", err)
		//return
	}
	r.C = GetConn(url)
}

func (r *Rich) Close() {
	err := r.C.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		logrus.Errorf("请求关闭失败:%v", err)
		return
	}
}

func (r *Rich) Ping() error {
	return r.C.WriteMessage(websocket.PingMessage, []byte(fmt.Sprintf("%d", time.Now().Unix())))
}

func (r *Rich) ReadAndPrint() {
	defer func() {
		if err := recover() ; err != nil {
			logrus.Warnf("recover 到panic err : %v",err)
			r.Signal <- true
			runtime.Goexit()
			time.Sleep(300 * time.Millisecond)
		}
	}()
	r.C.SetPingHandler(func(appData string) error {
		err := r.C.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(3*time.Second))
		if err != nil {
			logrus.Errorf("回写 ping 失败 ， err : %v", err)
			return err
		}

		return nil
	})
	r.C.SetPongHandler(func(appData string) error {
		logrus.Info("收到服务器回写pong")
		return nil
	})


	var temp int64
	for {
		select {
		case <- r.Down:
			logrus.Warnf("收到down信息 ,准备退出")
			r.Signal <- true
			runtime.Goexit()
			time.Sleep(300 * time.Millisecond)
		default:
			r.C.SetReadDeadline(time.Now().Add(10 * time.Second) )
			msgType, msg, err := r.C.ReadMessage()
			if err != nil {
				logrus.Errorf("读到消息错误：%s ", err.Error())
				r.Signal <- true
				runtime.Goexit()
				return
			}

			switch msgType {
			case websocket.TextMessage:
				e := &model.StreamText{}
				err := json.Unmarshal(msg, &e)
				if err != nil {
					logrus.Errorf(" unmarshal err : ", err)
				}

				str := ""
				if e.Data.IsSeller {
					str = "主动卖出"
				} else {
					str = "主动买入"
				}

				price, err := strconv.ParseFloat(e.Data.TradingPrice, 64)
				if err != nil {
					logrus.Errorf("ParseFloat fail , err : %v", err)
				}
				count, err := strconv.ParseFloat(e.Data.TradingCount, 64)
				if err != nil {
					logrus.Errorf("ParseFloat fail , err : %v", err)
				}

				temp = e.Data.Timestamp + 5000

				if time.Now().Unix() > temp {
					logrus.Warn("5秒未接收到数据 断线重连")
					r.Signal <- true
					runtime.Goexit()

				}


				logrus.Info(fmt.Sprintf("%s 现价为%.3f,数量为%.6f, 总值为%3f 的 %s",
					str, price, count, count*price, e.Data.TradingPair,
				))

				T = e.Data.Timestamp
			case websocket.BinaryMessage:
				logrus.Info("收到 BinaryMessage")
				logrus.Infof("读到BinaryMessage消息：%s", string(msg))

			}
		}



	}

}
