package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 2_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active:  true,
		},
	}))

	//Output:
	//100000
	//300000
	//0
	//0
}

func ExamplePaymentSources() {
	paymentSources := PaymentSources([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
			Name:    "Card1",
			PAN:     "5058 xxxx xxxx 7777",
		},
		{
			Balance: 1_000_00,
			Active:  false,
			Name:    "Card2",
			PAN:     "5058 xxxx xxxx 6666",
		},
		{
			Balance: -1_000_00,
			Active:  true,
			Name:    "Card3",
			PAN:     "5058 xxxx xxxx 8888",
		},
	})
	for _, paymentSource := range paymentSources {
		fmt.Println(paymentSource.Number)
	}

	//Output: 5058 xxxx xxxx 7777
}
