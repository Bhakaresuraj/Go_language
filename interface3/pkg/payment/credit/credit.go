package credit

import(
    "fmt"
)

type CreditCard struct{
    CardNumber string
    CSV string
    ExpiryDate string
}


func (c CreditCard)Pay(amount float64)string{

    msg := fmt.Sprintf(" Rs %.2f is debited successfully by credit card of  No : %s...!\n", amount,c.CardNumber)
	return msg
}
