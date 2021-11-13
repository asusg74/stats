package stats

import (
  "reflect"
  "testing"

  "github.com/asusg74/bank/v2/pkg/types"
)



func TestCategoriesAvg(t *testing.T) {
  payments := []types.Payment{
    {ID: 1, Category: "auto", Amount: 1_000},
    {ID: 2, Category: "auto", Amount: 2_000},
    {ID: 3, Category: "home", Amount: 3_000},
    {ID: 4, Category: "fun", Amount: 4_000},
  }
  expected := map[types.Category]types.Money{
    "auto" : 1500,
    "home" : 3000,
    "fun"  : 4000,
  }
  result := CategoriesAvg(payments)

  if !reflect.DeepEqual(expected, result) {
    t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
  }
}

