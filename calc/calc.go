package calc

import (
	"fmt"
	"math"
)

func PrincipalCalc(payment, periods, interest float64) float64 {
	return payment / ((interest * math.Pow(1+interest, periods)) / (math.Pow(1+interest, periods) - 1))
}

func PeriodsCalc(principal, payment, interest float64) float64 {
	return math.Ceil(math.Log(payment/(payment-interest*principal)) / math.Log(1+interest))
}

func PaymentCalc(principal, periods, interest float64) float64 {
	return principal * ((interest * math.Pow((1+interest), periods)) / (math.Pow((1+interest), periods) - 1))
}

func PaymentCalcDiff(principal, periods, interest, currMonth float64) float64 {

	return principal/periods + interest*(principal-(principal*(currMonth-1))/periods)
}
func PrintOverpayment(principal, sum float64) {
	fmt.Printf("\n%d\n", int(math.Ceil(sum)))
	fmt.Printf("Overpayment = %d", int(math.Ceil(sum)-principal))
}
