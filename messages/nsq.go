/*
 * Copyright (C) 2016 Miguel Moll
 *
 * This file is part of the Paranoid Minion
 *
 * The Paranoid Minion is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The Paranoid Minion is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with The Paranoid Minion.  If not, see <http://www.gnu.org/licenses/>.
 */

package messages

import (
	"dynastic.ninja/paranoid/minion"
	"github.com/nsqio/go-nsq"
)

type NsqDriver struct {
	Config          *nsq.Config
	Consumer        *nsq.Consumer
	Host            string
	MsgConverter    MessageConverter
	ReceiverChannel chan<- QueueData
}

func (n *NsqDriver) Initialize(host, queueName string) (err error) {

	n.MsgConverter = &MsgPackFormat{}

	n.Config = nsq.NewConfig()
	n.Host = host
	n.Consumer, err = nsq.NewConsumer(queueName, queueName, n.Config)
	if err != nil {
		return err
	}

	n.Consumer.SetLogger(minion.Log, nsq.LogLevelInfo)
	n.Consumer.AddHandler(n)

	return nil
}

func (n *NsqDriver) Connect() error {

	err := n.Consumer.ConnectToNSQLookupd(n.Host)
	if err != nil {
		minion.Log.Info("Host:", n.Host)
		return err
	}

	return nil
}

func (n *NsqDriver) SetReceiveChannel(rc chan<- QueueData) {
	n.ReceiverChannel = rc
}

func (n *NsqDriver) HandleMessage(message *nsq.Message) error {

	var data QueueData
	err := n.MsgConverter.ToQueueData(message.Body, &data)
	if err != nil {
		return err
	}

	n.ReceiverChannel <- data
	return nil
}

func (n *NsqDriver) Shutdown() {
	n.Consumer.Stop()
}

//func SendNSQ() {
//	config := nsq.NewConfig()
//	w, _ := nsq.NewProducer("127.0.0.1:4150", config)
//
//	err := w.Publish("write_test", []byte("test"))
//	if err != nil {
//		minion.Log.Error("Could not connect")
//	}
//
//	w.Stop()
//}
//
//func RecieveNSQ() {
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//
//	config := nsq.NewConfig()
//	q, _ := nsq.NewConsumer("write_test", "ch", config)
//	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
//		log.Printf("Got a message: %v", message)
//		wg.Done()
//		return nil
//	}))
//	err := q.ConnectToNSQD("127.0.0.1:4150")
//	if err != nil {
//		log.Panic("Could not connect")
//	}
//	wg.Wait()
//}
