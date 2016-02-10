package herder

import (
	"dynastic.ninja/paranoid/minion"
	"dynastic.ninja/paranoid/minion/transaction"
)

var TYPE = "patcher"

type Patcher struct {
}

func (p *Patcher) Type() string {
	return TYPE
}

func (p *Patcher) Run(t *transaction.Transaction, s *transaction.Supervisor) {

	minion.Log.Info("In a herder run's method!!")

}
