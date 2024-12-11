package models

type PaymentData struct {
	Pan      string `json:"pan"`      // 16-19 digits (Primary Account Number)
	Amount   string `json:"amount"`   // 12-digit string (Payment amount)
	Merchant string `json:"merchant"` // 15-digit string (Merchant ID)
	Terminal string `json:"terminal"` // 8-digit string (Terminal ID)
	Acquirer string `json:"acquirer"` // 6-9 digit string (Acquirer ID)
	Date     string `json:"data"`     // Date in YYYYMMDD format (e.g., "20241201")
	Time     string `json:"time"`     // Time in HHMMSS format (e.g., "121233")
	Trace    string `json:"trace"`    // 6-digit string (Trace ID)
	Rrn      string `json:"rrn"`      // 12-digit string (Retrieval Reference Number)
	Cvv2     string `json:"cvv2"`     // 3-4 digits (Card Verification Value)
	Exp      string `json:"exp"`      // Expiry date in MMYY format (e.g., "1122")
	Pin      string `json:"pin"`      // 6-digit string (PIN)
}
