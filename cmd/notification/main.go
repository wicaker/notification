package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/wicaker/notification/config"
	"github.com/wicaker/notification/internal/transport"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		}).Println(err)
	}
}

func main() {
	var (
		gomail     = config.NewGomailDialer()
		rabbitConn = config.NewRabbitmq()
		errChan    = make(chan error)
	)
	defer close(errChan)
	defer func() {
		err := rabbitConn.Close()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"at": time.Now().Format("2006-01-02 15:04:05"),
			}).Errorln(err)
		}
	}()

	go func() {
		for {
			transport.Handler(rabbitConn.Queue, gomail)

			// reconnect when rabbitmq server terminated accidentally
			err := <-rabbitConn.ErrorChannel
			if !rabbitConn.IsClose {
				rabbitConn.Reconnect(err)
			} else {
				errChan <- err
			}

		}
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	logrus.WithFields(logrus.Fields{
		"at": time.Now().Format("2006-01-02 15:04:05"),
	}).Errorln("terminated ", <-errChan)
}
