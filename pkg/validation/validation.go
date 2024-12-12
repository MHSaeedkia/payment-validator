package validation

import (
	"errors"
	"strconv"
	"time"

	"github.com/MHSaeedkia/blu-bank-validation/models"
)

type Payment models.PaymentData

func (p *Payment) Validate() error {
	// Validate Pan : 16-19 digits (Primary Account Number)
	if len(p.Pan) != 16 {
		return errors.New("Invalid Pan")
	}

	// Validate Amount : 12-digit string (Payment amount)
	if len(p.Amount) != 12 {
		return errors.New("Invalid Amount")
	}

	// Validate Merchant : 15-digit string (Merchant ID)
	if len(p.Merchant) != 15 {
		return errors.New("Invalid Merchant")
	}

	// Validate Terminal : 8-digit string (Terminal ID)
	if len(p.Terminal) != 8 {
		return errors.New("Invalid Terminal")
	}

	// Validate Acquirer : 6-9 digit string (Acquirer ID)
	if len(p.Acquirer) < 6 || len(p.Acquirer) > 9 {
		return errors.New("Invalid Acquirer")
	}

	// Validate Data : Date in YYYYMMDD format (e.g., "20241201")
	if _, err := time.Parse("20060102", p.Date); err != nil {
		return errors.New("Invalid Date")
	}

	// Validate Trace : 6-digit string (Trace ID)
	if len(p.Trace) != 6 {
		return errors.New("Invalid Trace")
	}

	// Validate Rrn : 12-digit string (Retrieval Reference Number)
	if len(p.Rrn) != 12 {
		return errors.New("Invalid Rrn")
	}

	// Validate Cvv2 : 3-4 digits (Card Verification Value)
	if len(p.Cvv2) != 3 {
		return errors.New("Invalid Cvv2")
	}

	// Validate Pin : 6-digit string (PIN)
	if len(p.Pin) != 6 {
		return errors.New("Invalid Pin")
	}

	// Validate Time : Time in HHMMSS format (e.g., "121233")
	err := timeValidator(p.Time)

	// Validate Exp : Expiry date in MMYY format (e.g., "1122")
	err = expValidator(p.Exp)

	return err
}

func timeValidator(time string) error {
	if len(time) != 6 {
		return errors.New("Invalid Time")
	} else if hour, err := strconv.Atoi(time[0:2]); err != nil || hour < 0 || hour > 24 {
		return errors.New("Invalid Time - hour")
	} else if min, err := strconv.Atoi(time[2:4]); err != nil || min > 60 {
		return errors.New("Invalid Time - miniute")
	} else if sec, err := strconv.Atoi(time[4:6]); err != nil || sec > 60 {
		return errors.New("Invalid Time - secund")
	}
	return nil
}

func expValidator(exp string) error {
	if len(exp) != 6 {
		return errors.New("Invalid Exp")
	} else if month, err := strconv.Atoi(exp[0:2]); err != nil || month > 12 {
		return errors.New("Invalid Exp - Month")
	} else if year, err := strconv.Atoi(exp[2:4]); err != nil || year > 100 {
		return errors.New("Invalid Exp - Year")
	}
	return nil
}
