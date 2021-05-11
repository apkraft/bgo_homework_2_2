package transfer

import (
	"testing"

	"github.com/apkraft/bgo_homework_2_1/pkg/card"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		TransferService *Service
	}

	type args struct {
		fromCard string
		toCard   string
		amount   int64
	}

	cardService := card.NewService("The Bank")

	insideTheBank := Fee{
		FeeInPercents:  0,
		MinFeeInCopeks: 0,
	}
	toAnotherBank := Fee{
		FeeInPercents:  5,
		MinFeeInCopeks: 10_00,
	}
	betweenOtherBanks := Fee{
		FeeInPercents:  15,
		MinFeeInCopeks: 30_00,
	}

	transferService := NewService(cardService, insideTheBank, toAnotherBank, betweenOtherBanks)

	cardService.NewCard(&card.Card{
		Issuer:   "Visa",
		Balance:  100_00,
		Currency: "RUB",
		Number:   "5106211234",
	})
	cardService.NewCard(&card.Card{
		Issuer:   "Visa",
		Balance:  100_00,
		Currency: "RUB",
		Number:   "5106215678",
	})
	cardService.NewCard(&card.Card{
		Issuer:   "Visa",
		Balance:  100_00,
		Currency: "RUB",
		Number:   "5106219012",
	})

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantError error
	}{
		{
			name: "Перевод внутри банка, денег достаточно",
			fields: fields{
				TransferService: transferService,
			},
			args: args{
				fromCard: "5106211234",
				toCard:   "5106215678",
				amount:   10_00,
			},
			wantTotal: 10_00,
			wantError: nil,
		},
		{
			name: "Перевод внутри банка, денег недостаточно",
			fields: fields{
				TransferService: transferService,
			},
			args: args{
				fromCard: "5106215678",
				toCard:   "5106219012",
				amount:   500_00,
			},
			wantTotal: 500_00,
			wantError: errNotEnoughMoney,
		},
		{
			name: "Перевод в другой банк, денег достаточно",
			fields: fields{
				TransferService: transferService,
			},
			args: args{
				fromCard: "5106215678",
				toCard:   "90129",
				amount:   10_00,
			},
			wantTotal: 20_00,
			wantError: errToCardNotFound,
		},
		{
			name: "Перевод в другой банк, денег недостаточно",
			fields: fields{
				TransferService: transferService,
			},
			args: args{
				fromCard: "5106215678",
				toCard:   "90129",
				amount:   500_00,
			},
			wantTotal: 525_00,
			wantError: errNotEnoughMoney,
		},
		{
			name: "Перевод из другого банка",
			fields: fields{
				TransferService: transferService,
			},
			args: args{
				fromCard: "56789",
				toCard:   "5106219012",
				amount:   500_00,
			},
			wantTotal: 575_00,
			wantError: errFromCardNotFound,
		},
		{
			name: "Перевод между другими банками",
			fields: fields{
				TransferService: transferService,
			},
			args: args{
				fromCard: "56789",
				toCard:   "90129",
				amount:   500_00,
			},
			wantTotal: 575_00,
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTotal, gotOk := transferService.Card2Card(tt.args.fromCard, tt.args.toCard, tt.args.amount)
			t.Log(gotTotal, tt.wantTotal, gotOk, tt.wantError)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantError {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantError)
			}
		})
	}
}
