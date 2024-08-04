package main

import (
	"LoanCalculator/calc"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
)

func main() {

	payment := flag.Float64("payment", 0, "Enter payment amount")
	principal := flag.Float64("principal", 0, "Enter loan principal amount")
	periods := flag.Float64("periods", 0, "Enter number of months to repay loan")
	interest := flag.Float64("interest", -1, "Enter interest rate")
	diffOrAnn := flag.String("type", "none", "Enter payment type")

	flag.Parse()

	if len(os.Args) != 5 ||
		int(*payment) < 0 ||
		int(*principal) < 0 ||
		int(*periods) < 0 ||
		int(*interest) < 0 ||
		(*diffOrAnn == "diff" && (int(*periods) < 0 || int(*principal) < 0)) ||
		(*diffOrAnn != "diff" && *diffOrAnn != "annuity") {
		err := errors.New("Incorrect parameters")
		fmt.Println(err)
	}

	*interest /= 1200
	sum := *payment * *periods

	if math.Round(*payment) == 0 {

		if *diffOrAnn == "annuity" {
			*payment = math.Ceil(calc.PaymentCalc(*principal, *periods, *interest))
			fmt.Printf("Your annuity payment = %d!\n", int(*payment))
			sum = *payment * *periods

		} else {
			for i := 1; i <= int(*periods); i++ {
				*payment = math.Ceil(calc.PaymentCalcDiff(*principal, *periods, *interest, float64(i)))
				fmt.Printf("Month %d: payment is %d\n", i, int(*payment))
				sum += *payment
			}
		}
		calc.PrintOverpayment(*principal, sum)
	}

	if math.Round(*principal) == 0 {
		*principal = math.Floor(calc.PrincipalCalc(*payment, *periods, *interest))
		fmt.Printf("Your loan principal = %d!", int(*principal))
		sum = *payment * *periods
		calc.PrintOverpayment(*principal, sum)
	}

	if math.Round(*periods) == 0 {
		*periods = calc.PeriodsCalc(*principal, *payment, *interest)

		var y int
		y = int(*periods / 12)
		var m int
		m = int(*periods) % 12
		var mStr string
		var yStr string

		switch m {
		case 0:
			mStr = ""
		case 1:
			mStr = "1 month"
		default:
			mStr = fmt.Sprintf("%d months", m)
		}
		switch y {
		case 0:
			yStr = ""
		case 1:
			yStr = "1 year "
		default:
			yStr = fmt.Sprintf("%d years", y)
		}
		if y > 0 && m > 0 {
			yStr += " and "
		}

		fmt.Printf("It will take %s%s to repay this loan!\n", yStr, mStr)

		sum = *payment * *periods

		calc.PrintOverpayment(*principal, sum)
	}
}
