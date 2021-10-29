package stats

import "github.com/asusg74/bank/pkg/types"

// Calculates avg sum of payments
func Avg(payments []types.Payment) types.Money {
	allSum := types.Money(0)
	allCount := types.Money(0)

	if len(payments) < 1 {
		return allSum
	}

	for _, payment := range payments {
			allSum += payment.Amount
			allCount += 1
	}

	if allCount < 1 {
		return allSum
	}

	avgSum := allSum / allCount

	return avgSum
}

// Calculate sum of payments in category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	allSum := types.Money(0)

	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		allSum += payment.Amount
	}

	return allSum
}