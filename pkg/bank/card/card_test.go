package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	fmt.Println(PaymentSource([]types.Card{
		{
			PAN: "5500 xxxx xxxx 3330",
			Balance: 1_000,
			Active:  true,
		},
		{
			PAN: "5500 xxxx xxxx 3332",
			Balance: 2_000,
			Active:  true,
		},
	}))

	// 	Output:
	// 950
}
