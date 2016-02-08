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

package transaction

import (
	"dynastic.ninja/paranoid/minion"
	"dynastic.ninja/paranoid/minion/msgqueue"
)

type Supervisor struct {
	OutgoingChannel chan<- msgqueue.QueueData
	IncomingChannel <-chan msgqueue.QueueData
}

// Run starts the supervisor to manage and route transactions.
// It's a blcoking call.
func (s *Supervisor) Run() {
	for {
		select {
		case data := <-s.IncomingChannel:
			minion.Log.Info("Incoming data: ", data)
		}
	}
}
