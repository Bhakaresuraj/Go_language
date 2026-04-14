package main

import (
    "fmt"
    "github.com/Bhakaresuraj/Go_language/interfaces-2/pkg/payments"
    "github.com/Bhakaresuraj/Go_language/interfaces-2/pkg/payments/upi"
)

func Checkout (method payments.PaymentMethod ,amount float64) string{
    
    msg:=method.Pay(amount)
    return msg
}




func main(){

    fmt.Println("Welcome to Interface practice ")
    shubhamUpi := upi.UPIPayment{UpiId :"suraj@oksbi" , App:"Gpay"}
    msg:=Checkout(shubhamUpi,24.3)
    fmt.Printf("Payment Successfull : %s\n",msg)
    DipakUpi := upi.UPIPayment{UpiId :"dipak@oksbi" , App:"PhonePay"}
    msg = Checkout(DipakUpi,24.3)
    fmt.Printf("Payment Successfull : %s\n",msg)


}
