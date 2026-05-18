package main
import(
    "fmt"
    "github.com/Bhakaresuraj/Go_language/interface3/pkg/payment"
    "github.com/Bhakaresuraj/Go_language/interface3/pkg/payment/upi"
)



func checkout(method payments.PaymentMethod ,amount float64 ) string {
    
}


func main(){
    fmt.Println("Payment interface example ")
    upi:=upi.Upipayment{upiId :"suraj@oksbi" , App :"Gpay"}
}
