/* Paranoid Minion
 * Copyright (C) 2016 Miguel Moll
 *
 * This software is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This software is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this software.  If not, see <http://www.gnu.org/licenses/>.
 */

package msgqueue

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/nsqio/go-nsq"
)

type NsqDriver struct {
	Config          *nsq.Config
	Consumer        *nsq.Consumer
	Host            string
	MsgConverter    MessageConverter
	IncomingChannel chan<- QueueData

	nsqLog *NsqLogger
}

func (n *NsqDriver) Initialize(host, queueName string) (err error) {

	n.MsgConverter = &MsgPackFormat{}

	n.Config = nsq.NewConfig()
	n.Host = host
	n.Consumer, err = nsq.NewConsumer(queueName, queueName, n.Config)
	if err != nil {
		return err
	}

	n.nsqLog = &NsqLogger{logrus.StandardLogger()}
	n.Consumer.SetLogger(n.nsqLog, nsq.LogLevelInfo)
	n.Consumer.AddHandler(n)

	return nil
}

func (n *NsqDriver) Connect() error {

	err := n.Consumer.ConnectToNSQLookupd(n.Host)
	if err != nil {
		return err
	}

	return nil
}

func (n *NsqDriver) SetIncomingChannel(rc chan<- QueueData) {
	n.IncomingChannel = rc
}

func (n *NsqDriver) HandleMessage(message *nsq.Message) error {

	var data QueueData
	err := n.MsgConverter.ToQueueData(message.Body, &data)
	if err != nil {
		return err
	}

	n.IncomingChannel <- data
	return nil
}

func (n *NsqDriver) Shutdown() {
	n.Consumer.Stop()
}

type NsqLogger struct {
	*logrus.Logger
}

// Output is used as a helper method to integrate with nsq logging system.
// Calldepth is ignored for now. This is a hack!!
func (n *NsqLogger) Output(calldepth int, s string) error {

	// Prefix being searched for is logging convention used by the go-nsq client
	if strings.HasPrefix(s, "INF") {
		n.Info(strings.TrimSpace(strings.Replace(s, "INF", "", 1)))
	} else if strings.HasPrefix(s, "DBG") {
		n.Debug(strings.TrimSpace(strings.Replace(s, "DBG", "", 1)))
	} else if strings.HasPrefix(s, "WRN") {
		n.Warn(strings.TrimSpace(strings.Replace(s, "WRN", "", 1)))
	} else if strings.HasPrefix(s, "ERR") {
		n.Error(strings.TrimSpace(strings.Replace(s, "ERR", "", 1)))
	}

	return nil
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
