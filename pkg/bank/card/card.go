package card

import (
	"Lesson7/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card:=types.Card{
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Active: true,
		Color: color,
		Name: name,
		MinBalance: 0,
	}
	return card
}

func Total(cards []types.Card) types.Money {

	var sum types.Money

	for _, bezh := range cards {
		sum += bezh.Balance
	}
	return sum
}
	