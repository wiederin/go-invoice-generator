package components

import (
	"errors"

	"github.com/shopspring/decimal"
)

var ErrInvalidDiscount = errors.New("invalid discount")

const (
	DiscountTypeAmount  string = "amount"
	DiscountTypePercent string = "percent"
)

type Discount struct {
	Percent string `json:"percent,omitempty"` 
	Amount  string `json:"amount,omitempty"`  

	_percent decimal.Decimal
	_amount  decimal.Decimal
}

func (d *Discount) Prepare() error {
	if len(d.Percent) == 0 && len(d.Amount) == 0 {
		return ErrInvalidDiscount
	}

	if len(d.Percent) > 0 {
		percent, err := decimal.NewFromString(d.Percent)
		if err != nil {
			return err
		}
		d._percent = percent
	}

	if len(d.Amount) > 0 {
		amount, err := decimal.NewFromString(d.Amount)
		if err != nil {
			return err
		}
		d._amount = amount
	}

	return nil
}

// getDiscount as return the discount type and value
func (t *Discount) getDiscount() (string, decimal.Decimal) {
	tax := "0"
	taxType := DiscountTypePercent

	if len(t.Percent) > 0 {
		tax = t.Percent
	}

	if len(t.Amount) > 0 {
		tax = t.Amount
		taxType = DiscountTypeAmount
	}

	decVal, _ := decimal.NewFromString(tax)

	return taxType, decVal
}
