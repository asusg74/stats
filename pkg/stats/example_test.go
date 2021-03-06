package stats

import (
	"fmt"

	"github.com/asusg74/bank/v2/pkg/types"
)

func ExampleAvg_positive() {
	fmt.Println(Avg([]types.Payment{{Amount: 10}, {Amount: 20}}))
	// Output:
	// 15
}

func ExampleAvg_null() {
	fmt.Println(Avg([]types.Payment{}))
	// Output:
	// 0
}

func ExampleTotalInCategory_positive() {
	fmt.Println(TotalInCategory([]types.Payment{{Amount: 10, Category: "Auto"},{Amount: 20, Category: "Internet"},}, "Auto"))
	// Output:
	// 10
}
