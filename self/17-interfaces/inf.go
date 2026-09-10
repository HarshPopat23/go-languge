package main

import "fmt"

type paymenter interface{
	pay(amount float32)
	refund(amount float32,account string)
}

type payment struct{
	geteway paymenter
}

func (p payment) makepayment(amount float32){
	// razorpaypaymenr := razorpay{}
	// stribepaypaymenr := stribe{}
	// stribepaypaymenr.pay(amount)
	p.geteway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32){
	fmt.Println("paying through razorpay", amount)
}

// type stribe struct{}

// func (s stribe) pay(amount float32){
// 	fmt.Println("paying through stripe", amount)
// }

type fakepayment struct{}

func (f fakepayment) pay(amount float32){
	fmt.Println("fake payment", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32){
	fmt.Println("paying through paypal", amount)
}

func (p paypal) refund(amount float32,account string){
	fmt.Println("refunding through paypal", amount,"to",account)
}

func main(){
	// stribepaypaymenr := stribe{}
	// razorpaypaymenr := razorpay{}
	paypalgew := paypal{}
	newpayment := payment{
		geteway: paypalgew,
	}
	newpayment.makepayment(100)
}