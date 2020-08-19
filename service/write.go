package service

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func WriteMessage(conn *websocket.Conn, msgType int, data []byte) {
	switch msgType {
	case websocket.TextMessage:
		go textMsg(conn,data)
	case websocket.BinaryMessage :
	case websocket.CloseMessage :
	case websocket.PingMessage :
	case websocket.PongMessage :
	}
	for {

		err := conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			logrus.Println("write:", err)
			return
		}
		logrus.Infof("读到消息类型：%d", msgType)
		//logrus.Infof("读到消息：%s", string(msg))
		return
	}
}

func textMsg (conn *websocket.Conn,data []byte){
	for {
		err := conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			logrus.Println("write:", err)
			return
		}
	}
}