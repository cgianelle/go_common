package common

import (
	"bufio"
	"fmt"	
)

func ScanLine(message string, scanner *bufio.Scanner) func() (string, error) {
	return func() (string, error) {
		fmt.Print(message)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
                        fmt.Println(err)
			return "", err
		} else {
        		return scanner.Text(), nil
		}
	}
}

