package filtrByCategory

import (
 "github.com/surush3005/Bank.git/pkg/bank/types"
)
	


func FiltrByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filetred []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filetred = append(filetred, payment)
		}
	}
}
