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

var herder = MockHerder{}

var outgoingChannel = make(chan<- msgqueue.QueueData)
var incomingChannel = make(<-chan msgqueue.QueueData)
var s *transaction.Supervisor

func TestMain(m *testing.M) {
	setup()

	results := m.Run()

	os.Exit(results)
}

func setup() {
	s = transaction.NewSupervisor(outgoingChannel, incomingChannel)
}

func TestSupervisorSetGet(t *testing.T) {

	s.RegisterHerder(herder)

	returned := s.Herder(herder.Type())

	if returned != herder {
		t.Error("Did not return the expected herder.",
			"Expected:", herder.Type(),
			"Got:", returned.Type(),
		)
	}
}

func TestToTransaction(t *testing.T) {

	id := "1234567"
	action := "doit"

	qd := msgqueue.QueueData{}
	qd["id"] = id
	qd["type"] = herder.Type()
	qd["action"] = action

	trans := &transaction.Transaction{}
	trans.Type = herder.Type()
	trans.Action = action
	trans.Id = id

	returned, _ := s.ToTransacton(qd)
	if trans.String() != returned.String() {
		t.Error("Did not return the expected herder.", "\n",
			"Expected:", trans, "\n",
			"Got:", returned,
		)
	}
}
