package stats

import (
	"fmt"

	"github.com/Romiz01/bank101/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100_00,
			Category: "Mobile",
		},
		{
			ID:       2,
			Amount:   5_00,
			Category: "Mobile",
		},
		{
			ID:       3,
			Amount:   10_00,
			Category: "Mobile",
		},
		{
			ID:       4,
			Amount:   10_00,
			Category: "Mobiles",
		},
	}
	avg := Avg(payments)
	fmt.Println(avg)

	// Output:
	//3125
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   5_000_00,
			Category: "Mobile",
		},
		{
			ID:       2,
			Amount:   10_000_00,
			Category: "Cars",
		},
	}

	CategoryTo := types.Category("Mobile")
	TotalC := TotalInCategory(payments, CategoryTo)
	fmt.Println(TotalC)
	//Output:  500000

}
