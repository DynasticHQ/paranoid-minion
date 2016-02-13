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

import "dynastic.ninja/paranoid/minion/msgqueue"

type Supervisor struct {
	OutgoingChannel chan<- msgqueue.QueueData
	IncomingChannel <-chan msgqueue.QueueData

	herders map[string]Herder
}

// NewSupervisor creates a transaction.Supervisor type
func NewSupervisor(outgoing chan<- msgqueue.QueueData, incoming <-chan msgqueue.QueueData) *Supervisor {

	s := &Supervisor{}
	s.OutgoingChannel = outgoing
	s.IncomingChannel = incoming
	s.herders = make(map[string]Herder)

	return s
}

// RegisterHerder registers a herder with the Supervior.
func (s *Supervisor) RegisterHerder(h Herder) {
	if s.herders == nil {
		s.herders = make(map[string]Herder)
	}
	s.herders[h.Type()] = h
}

// Herder returns the herderType specified
func (s *Supervisor) Herder(herderType string) Herder {
	return s.herders[herderType]
}

// Run starts the supervisor to manage and route transactions.
// It's a blcoking call.
func (s *Supervisor) Run() {

	for {

		data := <-s.IncomingChannel
		trans, _ := s.ToTransacton(data)

		s.Herder(trans.Type).Run(trans, s)
	}
}

// ToTransaction converts QueueData qd and returns a pointer to Transaction.
// Return an error if unable to convert.
func (s *Supervisor) ToTransacton(qd msgqueue.QueueData) (*Transaction, error) {

	t := &Transaction{}
	t.Type = qd["type"].(string)

	return t, nil
}
