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

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"dynastic.ninja/paranoid/minion"
	"dynastic.ninja/paranoid/minion/herder"
	"dynastic.ninja/paranoid/minion/internal"
	"dynastic.ninja/paranoid/minion/msgqueue"
	"dynastic.ninja/paranoid/minion/transaction"
)

var (
	configPath  = flag.String("config", "minion.toml", "TOML file used by the Minion.")
	showVersion = flag.Bool("version", false, "print version string")
)

func main() {

	flag.Parse()

	if *showVersion {
		fmt.Printf(version.String("Paranoid Minion"))
		return
	}

	err := minion.InitConfig(*configPath)
	if err != nil {
		fmt.Println("Unable to load config file. Error:", err)
		fmt.Println("Using minion defaults.")
	}

	minion.InitLogging(minion.Config)

	outgoingMsg := make(chan msgqueue.QueueData)
	incomingMsg := make(chan msgqueue.QueueData)

	supervisor := transaction.NewSupervisor(outgoingMsg, incomingMsg)
	RegisterHerders(supervisor)

	msgQueue := &msgqueue.NsqDriver{}
	msgQueue.SetIncomingChannel(incomingMsg)
	msgQueue.Initialize("127.0.0.1:4161", "agent_id")
	msgQueue.Connect()

	minion.Log.Info("Minion has started. Ready Up!")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go supervisor.Run()

	for {
		select {
		case <-sigChan:
			msgQueue.Shutdown()
			minion.Log.Info("Minion has stopped.")
			return
		}
	}
}

func RegisterHerders(s *transaction.Supervisor) {
	p := herder.Patcher{}
	s.RegisterHerder(&p)
}
