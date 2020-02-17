package main

import "testing"
import "time"

func TestAmount(t *testing.T) {
	var valueInteger int32 = 1234
	var valueDecimal float64 = 12.34
	var amount = Amount{valueInteger, valueDecimal}

	if valueInteger != amount.Integer() {
		t.Errorf("Amount integer was incorrect, got: %d, want: %d.", amount.Integer(), valueInteger)
	}

	if valueDecimal != amount.Decimal() {
		t.Errorf("Amount decimal was incorrect, got: %b, want: %b.", amount.Decimal(), valueDecimal)
	}
}

func TestProfessionalTaxTransferAmount(t *testing.T) {
	var amountUsd = Amount{1111, 11.11}
	var amountBrl = Amount{2222, 22.22}
	var amountEur = Amount{3333, 33.33}

	var p = ProfessionalTaxTransferAmount{amountUsd, amountBrl, amountEur}

	if amountUsd.Integer() != p.Usd().Integer() {
		t.Errorf("ProfessionalTaxTransferAmount Usd was incorrect")
	}

	if amountBrl.Integer() != p.Brl().Integer() {
		t.Errorf("ProfessionalTaxTransferAmount Brl was incorrect")
	}

	if amountEur.Integer() != p.Eur().Integer() {
		t.Errorf("ProfessionalTaxTransferAmount Eur was incorrect")
	}
}

func TestProfessionalTaxTransfer(t *testing.T) {
	var amountUsd = Amount{1111, 11.11}
	var amountBrl = Amount{2222, 22.22}
	var amountEur = Amount{3333, 33.33}

	time := time.Now()
	consultDate := time.Format("2006-01-02 15:04:05")
	description := "Foo Bar"
	p := ProfessionalTaxTransferAmount{amountUsd, amountBrl, amountEur}
	tax := ProfessionalTaxTransfer{consultDate, description, p}

	if consultDate != tax.ConsultDate() {
		t.Errorf("ProfessionalTaxTransfer consultDate was incorrect, got: %s, want: %s.", tax.ConsultDate(), consultDate)
	}

	if description != tax.Description() {
		t.Errorf("ProfessionalTaxTransfer description was incorrect, got: %s, want: %s.", tax.Description(), description)
	}

	if amountUsd.Integer() != tax.Amount().Usd().Integer() {
		t.Errorf("ProfessionalTaxTransfer Amount Usd was incorrect")
	}

	if amountBrl.Integer() != tax.Amount().Brl().Integer() {
		t.Errorf("ProfessionalTaxTransfer Amount Brl was incorrect")
	}

	if amountEur.Integer() != tax.Amount().Eur().Integer() {
		t.Errorf("ProfessionalTaxTransfer Amount Eur was incorrect")
	}
}