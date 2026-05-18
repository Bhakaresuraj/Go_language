package main

import (
	"fmt"
	"github.com/Bhakaresuraj/Go_language/interface3/pkg/payment"
	"github.com/Bhakaresuraj/Go_language/interface3/pkg/payment/upi"
	"github.com/Bhakaresuraj/Go_language/interface3/pkg/payment/credit"
)

func checkout(method payments.PaymentMethod, amount float64) string {

	msg := method.Pay(amount)
	return msg

}

func main() {
	fmt.Println("Payment interface example ")
	surajupi := upi.Upipayment{ UpiId : "suraj@oksbi", App: "Gpay"}
    surajcredit := credit.CreditCard{CardNumber:"1234567890",CSV :"3032",ExpiryDate:"12/10/2058"}
    msg:= checkout(surajupi, 40.10)
	fmt.Printf("Payment Successful : %s", msg)
    msg = checkout(surajcredit, 40.10)
	fmt.Printf("Payment Successful : %s", msg)
}
