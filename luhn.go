package main

import (
        "time"
        "os"
        "bufio"
        "fmt"
        "strconv"
)

func checkLuhn(cardNo string) (bool, time.Duration) {
        start := time.Now()
        nDigits := len(cardNo)

        nSum := 0
        isSecond := false

        for i := nDigits - 1; i >= 0; i-- {
                digit, _ := strconv.Atoi(string(cardNo[i]))

                if isSecond {
                        digit = digit * 2
                        if digit > 9 {
                                digit = digit - 9
                        }
                }

                nSum += digit
                isSecond = !isSecond
        }

        return nSum%10 == 0, time.Since(start)
}

func main() {
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Println("----- RAND LUHN's CHECKER -----")
        fmt.Print("Insert here the credit card number:  ")
        scanner.Scan()
        cardNo := scanner.Text()
        res, dur := checkLuhn(cardNo)
        if  res == true {
                fmt.Println("This is a valid card")
        } else {
                fmt.Println("This is not a valid card")
        }

        fmt.Println("Took ", dur," to check the credit card")
}
