package main

import (
	"fmt"

	"github.com/MHSaeedkia/blu-bank-validation/models"
	"github.com/MHSaeedkia/blu-bank-validation/pkg/ftp"
	"github.com/MHSaeedkia/blu-bank-validation/pkg/reporter"
)

const (
	FTP_SERVER = "localhost"
	FILE_NAME  = "/tmp/example.txt"
)

func main() {
	report := models.Report{
		Pan:      "6219861901234567",
		PrCode:   "000000",
		Amount:   "000000012345",
		Trace:    "269431",
		Ldate:    "1125",
		Ltime:    "090000",
		Termid:   "34567890",
		Acquirer: "00000603799",
		TempType: 59,
	}

	err := reporter.Reporter(report, FILE_NAME)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	err = ftp.FTPClient(FTP_SERVER, FILE_NAME, "example.txt")
	if err != nil {
		fmt.Errorf(err.Error())
	}
}
