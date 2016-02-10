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

package transaction_test

import (
	"os"
	"testing"

	"dynastic.ninja/paranoid/minion/msgqueue"
	"dynastic.ninja/paranoid/minion/transaction"
)

type MockHerder struct{}

func (m MockHerder) Type() string {
	return "MockHerder"
}

func (m MockHerder) Run(t *transaction.Transaction, s *transaction.Supervisor) {
	return
}

var outgoingChannel = make(chan<- msgqueue.QueueData)
var incomingChannel = make(<-chan msgqueue.QueueData)
var s *transaction.Supervisor

func TestMain(m *testing.M) {
	setup()

	results := m.Run()

	os.Exit(results)
}

func setup() {
	s = &transaction.Supervisor{}
	s.OutgoingChannel = outgoingChannel
	s.IncomingChannel = incomingChannel
}

func TestNewSupervisor(t *testing.T) {
	super := transaction.NewSupervisor(outgoingChannel, incomingChannel)

	// TODO(miguelmoll): Not sure this is a useful way to check supervisor equality.
	// Needs to be improved.
	if super.OutgoingChannel != s.OutgoingChannel &&
		super.IncomingChannel != s.IncomingChannel {
		t.Error("Supervisors are not equal.",
			"Expected:", s,
			"Got:", super,
		)
	}
}

func TestSupervisorSetGet(t *testing.T) {

	herder := MockHerder{}
	s.RegisterHerder(herder)

	returned := s.Herder(herder.Type())

	if returned != herder {
		t.Error("Did not return the expected herder.",
			"Expected:", herder.Type(),
			"Got:", returned.Type(),
		)
	}

}
