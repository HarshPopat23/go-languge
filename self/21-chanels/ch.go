package main

import "fmt"

// func process(numchan chan int) {

// 	for num := range numchan {
// 		fmt.Println("Processing numchan", num)
// 		time.Sleep(time.Millisecond * 500)
// 	}

// }

// func sum(result chan int,num1 int,num2 int){
// 	numresult := num1+num2
// 	result <- numresult
// }

// func task(done chan bool){

// 	defer func(){done<-true}()

// 	fmt.Println("Prosecing...")
// }

// func emailsender(emailchan chan string,done chan bool){
// 	defer func() {done<-true}()

// 	for email := range emailchan{
// 		fmt.Println("sending email to",email)
// 		time.Sleep(time.Microsecond*500)
// 	}
// }

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func(){
		chan1<-10
	}()

	go func(){
		chan2<-"Harsh"
	}()

	for i:=0;i<2;i++{
		select{
		case chan1val:=<-chan1:
			fmt.Println("Receive from chan1",chan1val)
		case chan2val:=<-chan2:
			fmt.Println("Receive from chan2",chan2val)
		}
	}

	// emailchan := make(chan string,100)

	// done := make(chan bool)

	// go emailsender(emailchan,done)

	// for i:=0;i<5;i++{
	// 	emailchan <- fmt.Sprintf("%d@gmail.com",i)
	// }

	// close(emailchan)

	// emailchan <-"da@gmail.com"
	// emailchan <-"dae@gmail.com"

	// fmt.Println(<-emailchan)
	// fmt.Println(<-emailchan)

	// <- done

	// done := make(chan bool)

	// go task(done)

	// <-done

	// result := make(chan int)

	// go sum(result,4,5)

	// res:= <-result

	// fmt.Println(res)

	// numchan := make(chan int)

	// go process(numchan)

	// for{
	// 	numchan<-rand.Intn(100)
	// }

	// messagechan := make(chan string)

	// messagechan <- "Ping"

	// msg := <-messagechan

	// fmt.Println(msg)

}