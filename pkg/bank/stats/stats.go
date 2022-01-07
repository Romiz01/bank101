package stats

import (
	"github.com/Romiz01/bank101/bank/pkg/bank/types"
	// "bank/pkg/bank/types"
)

// Avg 1
func Avg(payments []types.Payment) types.Money {
	cPayments := types.Money(len(payments))
	sPaymenys := types.Money(0)
	for _, payment := range payments {
		mPayments := payment.Amount
		sPaymenys += mPayments
	}
	return sPaymenys / cPayments
}

// TotalInCategory 2
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	PaySum := types.Money(0)
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		PayMoney := payment.Amount
		PaySum += PayMoney

	}
	return PaySum
}
