package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Validator struct {
	field string
	value string
}

func Field(field string) *Validator {
	return &Validator{
		field: strings.ToLower(field),
	}
}

func (v *Validator) Value(value string) error {
	v.value = value
	switch v.field {
	case "pan":
		return panValidator(value)
	case "amount":
		return amountValidator(value)
	case "merchant":
		return merchantValidator(value)
	case "terminal":
		return terminalValidator(value)
	case "acquirer":
		return acquirerValidator(value)
	case "date":
		return dateValidator(value)
	case "trace":
		return traceValidator(value)
	case "rrn":
		return rrnValidator(value)
	case "cvv2":
		return cvv2Validator(value)
	case "pin":
		return pinValidator(value)
	case "time":
		return timeValidator(value)
	case "exp":
		return expValidator(value)
	default:
		return fmt.Errorf("unknown field: %s", v.field)
	}
}

func luhnValidator(pan string) error {
	sum := 0
	for i := 0; i < len(pan); i++ {
		lun, _ := strconv.Atoi(string(pan[i]))
		if i%2 == 0 {
			if lun*2 > 9 {
				sum += ((lun * 2) % 10) + ((lun * 2) / 10)
			} else {
				sum += lun * 2
			}
		} else {
			sum += lun
		}
	}
	if sum%10 != 0 {
		return errors.New("invalid pan")
	}
	return nil
}

func digitValidator(field, name string) error {
	for _, char := range field {
		if !unicode.IsDigit(char) {
			return fmt.Errorf("invalid %s : must be digit all characters", name)
		}
	}
	return nil
}

func panValidator(pan string) error {
	if len(pan) != 16 {
		return errors.New("invalid pan : must be 16 digits")
	}

	err := digitValidator(pan, "pan")
	if err != nil {
		return err
	}

	return nil
}

func amountValidator(amount string) error {
	d, err := strconv.Atoi(amount)
	if err != nil {
		return fmt.Errorf("invalid amount")
	}
	if d < 1000 || d > 100000000000 {
		return fmt.Errorf("invalid amount bound")
	}
	return nil
}

func merchantValidator(merchant string) error {
	if len(merchant) > 15 {
		return errors.New("invalid merchant : must be at last 15 digits")
	}
	return digitValidator(merchant, "merchant")
}

func terminalValidator(terminal string) error {
	if len(terminal) != 8 {
		return errors.New("invalid terminal : must be 8 digits")
	}
	return digitValidator(terminal, "terminal")
}

func acquirerValidator(acquirer string) error {
	if len(acquirer) < 6 || len(acquirer) > 11 {
		return errors.New("Invalid acquirer configuration")
	}
	return digitValidator(acquirer, "acquirer")
}

func dateValidator(date string) error {
	dt, err := time.Parse("20060102", date)
	if err != nil {
		return errors.New("invalid date")
	}
	if dt.After(time.Now()) {
		return errors.New("date cannot be in the future")
	}
	return digitValidator(date, "date")
}

func traceValidator(trace string) error {
	if len(trace) != 6 {
		return errors.New("invalid trace : must be 6 digits")
	}
	return digitValidator(trace, "trace")
}

func rrnValidator(rrn string) error {
	if len(rrn) != 12 {
		return errors.New("invalid rrn : must be 12 digits")
	}
	return digitValidator(rrn, "rrn")
}

func cvv2Validator(cvv2 string) error {
	if len(cvv2) != 3 {
		return errors.New("invalid cvv2 : must be 3 digits")
	}
	return digitValidator(cvv2, "cvv2")
}

func pinValidator(pin string) error {
	if len(pin) != 6 {
		return errors.New("invalid pin : must be 6 digits")
	}
	return digitValidator(pin, "pin")
}

func timeValidator(time string) error {
	err := digitValidator(time, "time")
	if err != nil {
		return err
	}
	if len(time) != 6 {
		return errors.New("invalid time : must be in HHMMSS format")
	} else if hour, err := strconv.Atoi(time[0:2]); err != nil || hour < 0 || hour > 24 {
		return errors.New("invalid time - hour")
	} else if min, err := strconv.Atoi(time[2:4]); err != nil || min > 60 {
		return errors.New("invalid time - miniute")
	} else if sec, err := strconv.Atoi(time[4:6]); err != nil || sec > 60 {
		return errors.New("invalid time - secund")
	}
	return nil
}

func expValidator(exp string) error {
	err := digitValidator(exp, "exp")
	if err != nil {
		return err
	}
	if len(exp) != 4 {
		return errors.New("invalid exp : must be in MMYY format")
	} else if month, err := strconv.Atoi(exp[0:2]); err != nil || month > 12 {
		return errors.New("invalid exp - month")
	} else if year, err := strconv.Atoi(exp[2:4]); err != nil || year > 100 {
		return errors.New("invalid exp - year")
	}
	return nil
}
