package upi 
 
import (
    "fmt"
)

type UPIPayment struct {
    UpiId string
    App string
} 

func (u UPIPayment)Pay(amount float64) string{

    msg:=fmt.Sprintf("UPI Payment of %.2f successfully completed using %s through %s .\n ",amount ,u.UpiId,u.App )
    return msg
}
