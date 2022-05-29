package card

import (
	"bank/pkg/bank/types"
)

func Total(cards []types.Card) types.Money{
	sum := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}

		if card.Balance <= 0{
			continue
		}

		sum += card.Balance
	}
	return sum
}

func PaymentSource(cards []types.Card)types.Money{
	Pokupka := types.Money(50)
	pan := types.PAN("5500 xxxx xxxx 3330")
	for _, card := range  cards{
		if card.Active && card.Balance >= Pokupka && card.PAN == pan{
			Pokupka = card.Balance - Pokupka
			return Pokupka
		}else{
			return card.Balance
		}	
	}
	return 0 
}