package ftp

import (
	"fmt"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

// FTP server details
const (
	FTP_USER     = "admin"
	FTP_PASSWORD = "password"
)

func FTPClient(ftpServer, filePath, fileName string) error {
	// Connect to FTP server
	c, err := ftp.Dial(fmt.Sprintf("%s:%s", ftpServer, "2121"), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		fmt.Println("Dial error : ", err)
		return err
	}
	defer c.Quit()

	// Login to FTP server
	if err := c.Login(FTP_USER, FTP_PASSWORD); err != nil {
		fmt.Println("Login error : ", err)
		return err
	}

	fmt.Println("Successfully logged into the FTP server")

	// Open the local file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open error : ", err)
		return err
	}
	defer file.Close()

	// Upload the file to the FTP server
	if err := c.Stor(fileName, file); err != nil {
		fmt.Println("Store error : ", err)
		return err
	}

	fmt.Printf("Successfully uploaded %s to %s\n", filePath, ftpServer)
	return nil
}
