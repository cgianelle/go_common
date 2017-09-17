package common

import (
	"bufio"
	"os"
	"fmt"	
)

type EFP_Scanner struct {
	scanner *bufio.Scanner
}

//Constructor of sorts
func New_EFP_Scanner() *EFP_Scanner {
	efp := new(EFP_Scanner)
	efp.scanner = bufio.NewScanner(os.Stdin)
	return efp
}

//Create a scanner function with a customer message
func (efp *EFP_Scanner) CreateScanner(message string) func() (string, error) {
	return func() (string, error) {
		fmt.Print(message)
		efp.scanner.Scan()
		if err := efp.scanner.Err(); err != nil {
                        fmt.Println(err)
			return "", err
		} else {
        		return efp.scanner.Text(), nil
		}
	}
}
