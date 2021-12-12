package entity

import (
	"context"
	"log"
	"regexp"
	"time"
)

const _regexValidateCreditCardNumber = `^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`

// CreditCard é a entidade padrão de um cratão de credito
type CreditCard struct {
	number          string
	name            string
	expirationMonth int
	expirationYear  int
	cvv             int
}

// NewCreditCard cria um novo catão de credito, caso os paramentros não seja válidos, um erro é retornado
func NewCrediCard(number string, name string, expirationMonth int, expirationYear int, cvv int) (*CreditCard, error) {
	new := &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		cvv:             cvv,
	}

	if err := new.IsValid(); err != nil {
		log.Printf("context: %v, err: %s", context.Background(), ErrInvalidCreditCard.Error())
		return nil, err
	}

	return new, nil
}

// IsValid realiza a validação de uma entidade CreditCard
func (c CreditCard) IsValid() error {
	if err := c.validadeNumber(); err != nil {
		return err
	}

	if err := c.validadeMonth(); err != nil {
		return err
	}

	if err := c.validadeYear(); err != nil {
		return err
	}

	return nil
}

// validateNumber realiza a validação do numero do cartão de credito
func (c CreditCard) validadeNumber() error {
	if !regexp.MustCompile(_regexValidateCreditCardNumber).MatchString(c.number) {
		log.Printf("context: %v, err: %s", context.Background(), ErrInvalidCreditCardNumber.Error())
		return ErrInvalidCreditCardNumber
	}

	return nil
}

// validateMonth realiza a validação do mês do cartão de credito
func (c CreditCard) validadeMonth() error {
	if c.expirationMonth < 1 && c.expirationMonth > 13 {
		log.Printf("context: %v, err: %s", context.Background(), ErrInvalidCreditCardMonth.Error())
		return ErrInvalidCreditCardMonth
	}

	return nil
}

// validateYear realiza a validação do ano do cartão de credito
func (c CreditCard) validadeYear() error {
	if c.expirationYear < time.Now().Year() {
		log.Printf("context: %v, err: %s", context.Background(), ErrInvalidCreditCardYear.Error())
		return ErrInvalidCreditCardYear
	}

	return nil
}
