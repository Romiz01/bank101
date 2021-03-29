package stats

import (
	"github.com/Romiz01/bank101/pkg/bank/types"
	// "bank/pkg/bank/types"
)

// Avg
func Avg(payments []types.Payment) types.Money {
	cPayments := types.Money(len(payments))
	sPaymenys := types.Money(0)
	for _, payment := range payments {
		mPayments := payment.Amount
		sPaymenys += mPayments
	}
	return sPaymenys / cPayments
}
