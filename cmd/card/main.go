package main

import (
	"fmt"

	"github.com/apkraft/bgo_homework_2_1/pkg/card"
	"github.com/apkraft/bgo_homework_2_1/pkg/transfer"
)

func main() {

	cardSvc := card.NewService("The Bank")

	insideTheBank := transfer.Fee{
		FeeInPercents:  0,
		MinFeeInCopeks: 0,
	}
	toAnotherBank := transfer.Fee{
		FeeInPercents:  5,
		MinFeeInCopeks: 10_00,
	}
	betweenOtherBanks := transfer.Fee{
		FeeInPercents:  15,
		MinFeeInCopeks: 30_00,
	}
	transferSvc := transfer.NewService(cardSvc, insideTheBank, toAnotherBank, betweenOtherBanks)

	cardSvc.NewCard(&card.Card{
		Issuer:   "510621",
		Balance:  100_00,
		Currency: "RUB",
		Number:   "51062115",
	})
	cardSvc.NewCard(&card.Card{
		Issuer:   "510621",
		Balance:  100_00,
		Currency: "RUB",
		Number:   "51062123",
	})

	amount, status := transferSvc.Card2Card("51062115", "51062123", 10_00)
	fmt.Printf("Внутри банка: сумма с комиссией - %v, статус - %v\n\n", amount, status)

	amount, status = transferSvc.Card2Card("51062115", "56789", 10_00)
	fmt.Printf("В другой банк: сумма с комиссией - %v, статус - %v\n\n", amount, status)

	amount, status = transferSvc.Card2Card("12349", "56789", 10_00)
	fmt.Printf("Между другими банками: сумма с комиссией - %v, статус - %v\n\n", amount, status)

	amount, status = transferSvc.Card2Card("51062115", "51062123", 500_00)
	fmt.Printf("Недостаточно средств для перевода внутри банка: сумма с комиссией - %v, статус - %v\n\n", amount, status)

	amount, status = transferSvc.Card2Card("51062115", "56789", 500_00)
	fmt.Printf("Недостаточно средств для перевода в другой банк: сумма с комиссией - %v, статус - %v\n\n", amount, status)
}
