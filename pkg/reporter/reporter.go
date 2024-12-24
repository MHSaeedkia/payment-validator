package reporter

import (
	"fmt"
	"os"

	"github.com/MHSaeedkia/blu-bank-validation/models"
)

func Reporter(report models.Report, fileName string) error {
	rep := fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%v",
		report.Pan, report.PrCode,
		report.Amount, report.Trace,
		report.Ldate, report.Ltime,
		report.Termid, report.Acquirer,
		report.TempType)

	// Open the file for writing (create it if it doesn't exist, or truncate it if it does)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	// Ensure the file is closed at the end
	defer file.Close()

	// Content to write into the file

	// Write content to the file
	_, err = file.WriteString(rep)
	if err != nil {
		return err
	}

	fmt.Println("File written successfully:", fileName)
	return nil
}
