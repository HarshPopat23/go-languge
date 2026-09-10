package main

import "fmt"

type orderstatus string

const (
	Recived orderstatus = "Recived"
	Confiremed = "Confiremed"
	Prepared = "Prepared"
	Delivered = "Delivered"
)

func changeorderstatus(status orderstatus) {
	fmt.Println("Changeing status to", status)
}

func main() {
	changeorderstatus(Confiremed)
}						