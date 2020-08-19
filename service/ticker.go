package service

import (
	"github.com/sirupsen/logrus"
	"time"
)

func (rich *Rich) Ticker(url string) {
	ticker := time.NewTicker(time.Minute)
	t := time.NewTicker(time.Second * 5)
	defer t.Stop()
	defer ticker.Stop()
	for {
		select {
		case <-t.C:
			time.Sleep(time.Second)
			temp := time.Since(time.Unix(T, 0))
			if temp > time.Second*6 {
				rich.Down <- true
			}

		case <-ticker.C:
			logrus.Info("ping================================\n\n\n")
			err := rich.Ping()
			if err != nil {
				logrus.Errorf("[main] ping 失败 , err : %v ", err)
				rich.C = GetConn(url)
				if err != nil {
					logrus.Fatal("dial:", err)
				}
			}
		}
	}
}
