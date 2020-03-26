package main

import "fmt"

func main() {

var decimal_num int

fmt.Println (" This program converts decimal numbers to binary.\n Input a number you would like to convert (Integer numbers only!) : ")

fmt.Scan(&decimal_num)

fmt.Printf("The binary representation of that number is:" + " %b", decimal_num)

}
