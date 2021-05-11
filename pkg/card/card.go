package card

import (
	"strconv"
	"strings"
)

type Card struct {
	Id       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

type Service struct {
	BankName string
	Cards    []*Card
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) NewCard(card *Card) {
	card.Issuer = s.BankName
	s.Cards = append(s.Cards, card)
}

func (s *Service) FindCardByNumber(number string) (card *Card) {
	for _, card := range s.Cards {
		if number == card.Number && strings.HasPrefix(number, card.Issuer) {
			return card
		}
	}
	return nil
}

func (s *Service) LunaCardNumberCheck(number string) bool {
	reversedCardNumber := reverseCardNumber(number)
	cardNumberDigits := strings.Split(strings.ReplaceAll(reversedCardNumber, " ", ""), "")
	sum := 0

	for digitPlace := range cardNumberDigits {
		if digit, e := strconv.Atoi(cardNumberDigits[digitPlace]); e == nil {
			if (digitPlace+1)%2 != 0 {
				digit = digit * 2
				if digit > 9 {
					digit = digit - 9
				}
			}
			sum += digit
		} else {
			return false
		}
	}
	return sum%10 == 0
}

func reverseCardNumber(cardNumber string) string {
	runes := []rune(cardNumber)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
