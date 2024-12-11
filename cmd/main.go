package main

import (
	"fmt"

	"github.com/MHSaeedkia/blu-bank-validation/pkg/validation"
)

func main() {
	data := &validation.Payment{
		Pan:      "1234567890123456",
		Amount:   "100000000000",
		Merchant: "123456789012345",
		Terminal: "12345678",
		Acquirer: "123456",
		Date:     "20241201",
		Time:     "124233",
		Trace:    "123456",
		Rrn:      "123456789012",
		Cvv2:     "123",
		Exp:      "1122",
		Pin:      "123456",
	}
	if err := data.Validate(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("valid")
	}
}
