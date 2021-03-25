package card

import (
	"bank/pkg/bank/types"
)

//Total вычисляет общую сумму средств на всех картах
//Отрицательные суммы и суммы на заблакированных картах игнорируются
func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, operation := range cards {
		if !operation.Active {
			continue
		}
		if operation.Balance <= 0 {
			continue
		}

		sum += operation.Balance
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paymentSources []types.PaymentSource
	for _, operation := range cards {
		if !operation.Active {
			continue
		}
		if operation.Balance <= 0 {
			continue
		}

		paymentSources = append(paymentSources, types.PaymentSource{
			Type:    operation.Name,
			Number:  string(operation.PAN),
			Balance: operation.Balance,
		})
	}
	return paymentSources
}
