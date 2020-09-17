package card_test

import (
	"Lesson7/pkg/bank/card"
	"fmt"
	"Lesson7/pkg/bank/types"
	
)

func ExampleTotal()  {
	cards:=[]types.Card{
		{
			Balance: 1_000_00,
		    Active: true,
		},
	}
	fmt.Println(card.Total(cards))
	// Output: 100000
}