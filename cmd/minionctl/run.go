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

package main

import (
	"os"
	"os/signal"
	"syscall"

	"dynastic.ninja/paranoid/minion/herder"
	"dynastic.ninja/paranoid/minion/msgqueue"
	"dynastic.ninja/paranoid/minion/transaction"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

// runCmd is the primary command for the minion to start listening for messages.
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the paranoid minion",
	Long:  `Run the minion and listen for server messages.`,
	RunE:  Run,
}

func Run(cmd *cobra.Command, args []string) error {

	InitializeConfig()

	outgoingMsg := make(chan msgqueue.QueueData)
	incomingMsg := make(chan msgqueue.QueueData)

	supervisor := transaction.NewSupervisor(outgoingMsg, incomingMsg)
	registerHerders(supervisor)

	msgQueue := &msgqueue.NsqDriver{}
	msgQueue.SetIncomingChannel(incomingMsg)
	msgQueue.Initialize("127.0.0.1:4161", "agent_id")
	msgQueue.Connect()

	logrus.Info("Minion has started. Ready Up!")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go supervisor.Run()

	for {
		select {
		case <-sigChan:
			msgQueue.Shutdown()
			logrus.Info("Minion has stopped.")
			return nil
		}
	}

	return nil
}

func registerHerders(s *transaction.Supervisor) {
	p := herder.Patcher{}
	s.RegisterHerder(&p)
}
