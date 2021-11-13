package stats
import(
	"github.com/asusg74/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	allSum := types.Money(0)

	allCount := types.Money(0)

	if len(payments) < 1 {
		return allSum
	}

	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			allSum += payment.Amount
			allCount += 1
		}
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
		if payment.Status == types.StatusFail {
			continue
		}
		allSum += payment.Amount
	}

	return allSum
}
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {

	//Ответ 
	result := map[types.Category]types.Money{}
	
	//Количество платежей определенного типа
	count := map[types.Category]int{}
	for _, payment := range payments {
		result[payment.Category] += payment.Amount
		count[payment.Category]++
	}
	
	for category, cnt := range count {
		if cnt <= 0 {
			continue
		}
		result[category] /= types.Money(cnt)
	}
	return result
	}