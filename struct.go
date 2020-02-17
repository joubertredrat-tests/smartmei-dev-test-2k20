package main

import (
	"math"
)

type Amount struct {
	integer int32
	decimal float64
}

func (a *Amount) Integer() int32 {
	return a.integer
}

func (a *Amount) Decimal() float64 {
	return a.decimal
}

type ProfessionalTaxTransferAmount struct {
	usd Amount
	brl Amount
	eur Amount
}

func (t *ProfessionalTaxTransferAmount) Usd() *Amount {
	return &t.usd
}

func (t *ProfessionalTaxTransferAmount) Brl() *Amount {
	return &t.brl
}

func (t *ProfessionalTaxTransferAmount) Eur() *Amount {
	return &t.eur
}

type ProfessionalTaxTransfer struct {
	consultDate string
	description string
	amount ProfessionalTaxTransferAmount
}

func (p *ProfessionalTaxTransfer) ConsultDate() string {
	return p.consultDate
}

func (p *ProfessionalTaxTransfer) Description() string {
	return p.description
}

func (p *ProfessionalTaxTransfer) Amount() *ProfessionalTaxTransferAmount {
	return &p.amount
}

func createAmount(baseTax float64, rateCurrency float64) Amount {
	var floatAmount = math.Round((baseTax * rateCurrency) * 100) / 100
	var integerAmount = int32(floatAmount * 100)
	var amount = Amount{integerAmount, floatAmount}
	return amount
}