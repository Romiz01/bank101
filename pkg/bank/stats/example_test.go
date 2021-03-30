package stats

import (
	"fmt"

	// " bank/pkg/bank/types"
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
			Category: "Mobile",
		},
	}
	avg := Avg(payments)
	fmt.Println(avg)

	// Output:
	//3125
}
