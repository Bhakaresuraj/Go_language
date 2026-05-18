package upi

import (
	"fmt"
)

type Upipayment struct {
	UpiId string
	App   string
}

func (u Upipayment) Pay(amount float64) string {

	// logic to deduct amount from your upi account ..
	msg := fmt.Sprintf(" Rs %.2f is debited successfully from %s...!\n", amount, u.UpiId)
	return msg
}
