package card

import (
	"github.com/surush3005/Bank.git/pkg/bank/types"
)

func Issue(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		Id:       1000,
		PAN:      "5058 XXXX XXXX 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
}
