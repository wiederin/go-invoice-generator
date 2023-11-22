package components

import (
	"errors"

	"github.com/shopspring/decimal"
)

var ErrInvalidTax = errors.New("invalid tax")

const (
	TaxTypeAmount  string = "amount"
	TaxTypePercent string = "percent"
)

type Tax struct {
	Percent string `json:"percent,omitempty"` 
	Amount  string `json:"amount,omitempty"` 

	_percent decimal.Decimal
	_amount  decimal.Decimal
}

func (t *Tax) Prepare() error {
	if len(t.Percent) == 0 && len(t.Amount) == 0 {
		return ErrInvalidTax
	}

	if len(t.Percent) > 0 {
		percent, err := decimal.NewFromString(t.Percent)
		if err != nil {
			return err
		}
		t._percent = percent
	}

	if len(t.Amount) > 0 {
		amount, err := decimal.NewFromString(t.Amount)
		if err != nil {
			return err
		}
		t._amount = amount
	}

	return nil
}

func (t *Tax) getTax() (string, decimal.Decimal) {
	tax := "0"
	taxType := TaxTypePercent

	if len(t.Percent) > 0 {
		tax = t.Percent
	}

	if len(t.Amount) > 0 {
		tax = t.Amount
		taxType = TaxTypeAmount
	}

	decVal, _ := decimal.NewFromString(tax)

	return taxType, decVal
}
