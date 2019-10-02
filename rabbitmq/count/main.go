package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// docker run --rm -d --name=hi-rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
func main() {
	ch, _, err := ch()
	if nil != err {
		panic(err)
	} else {
		defer ch.Close()
	}

	go func() {
		for {
			qqq, err := ch.QueueDeclare("q1", false, false, false, false, nil)
			if nil != err {
				panic(err)
			} else {
				logrus.
					WithField("name", qqq.Name).
					WithField("consumers", qqq.Consumers).
					WithField("messages", qqq.Messages).
					Infoln("stats")
			}
			
			time.Sleep(1 * time.Second)
		}
	}()

	_, err = ch.QueueDeclare("q1", false, false, false, false, nil)
	if nil != err {
		panic(err)
	} else {
		_ = ch.QueueBind("q1", "test", "events", true, nil)

		go func() {
			for {
				err := ch.Publish("events", "test", false, false, amqp.Publishing{
					Body: []byte("testing only"),
				})

				if nil != err {
					panic(err)
				}

				time.Sleep(200 * time.Millisecond)
			}
		}()
	}

	time.Sleep(5 * time.Minute)
}

func ch() (*amqp.Channel, chan *amqp.Error, error) {
	if con, err := con(); nil != err {
		return nil, nil, err
	} else {
		if ch, err := con.Channel(); nil != err {
			return nil, nil, err
		} else {
			err = ch.ExchangeDeclare("events", "topic", false, false, false, false, nil)
			if nil != err {
				return nil, nil, err
			}

			return ch, con.NotifyClose(make(chan *amqp.Error)), nil
		}
	}
}

func con() (*amqp.Connection, error) {
	con, err := amqp.Dial("amqp://localhost:5672")
	if nil != err {
		return nil, err
	}

	return con, nil
}
